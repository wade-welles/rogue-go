#include "glow.h"
#include "engine.h"
#include "offsets.h"

#include <string>

typedef struct GlowObjectDefinition_t glow_t;

void Glow::Radar(uintptr_t entPtr, int lTeam, int rTeam)
{
    if (rTeam == TEAM_T || rTeam == TEAM_CT) {
        if (rTeam != lTeam) {
            m_mem.Write(entPtr + Netvar::CBasePlayer::m_bSpotted, true);
        }
    }
}

void Glow::Run()
{
    const bool Radar       = true;

    CGlowObjectManager manager;

    Log("[rogue] Starting Radar hack");

    while (!ShouldStop()) {
        struct iovec g_local[1024];
        struct iovec g_remote[1024];
        glow_t g_glow[1024];

        memset(g_local, 0, sizeof(g_local));
        memset(g_remote, 0, sizeof(g_remote));
        memset(g_glow, 0, sizeof(g_glow));

        if (!m_mem.Read(Offset::Client::GlowObjectManager, &manager)) {
            LogWait("[rogue] Failed to read GlowObjectManager");
            continue;
        }

        const size_t count = manager.Count();
        const uintptr_t GlowData = manager.Data();

        if (!m_mem.Read(GlowData, g_glow, sizeof(glow_t) * count)) {
            LogWait("[rogue] Failed to read m_GlowObjectDefinitions");
            continue;
        }

        uintptr_t localPlayer;
        if (!m_mem.Read(Offset::Client::LocalPlayer, &localPlayer)) {
            LogWait("[rogue] Failed to read local player address");
            continue;
        }

        int myTeam;
        if (!m_mem.Read(localPlayer + Netvar::CBaseEntity::m_iTeamNum, &myTeam)) {
            Wait();
            continue;
        }

        size_t writeCount = 0;

	Vector vecEyes;
	if (!m_mem.Read(localPlayer + Netvar::CBaseEntity::m_vecOrigin, &vecEyes)) {
            LogWait("[rouge] Failed to read CBaseEntity::m_vecOrigin");
            return;
	}

        for (size_t i = 0; i < count; ++i) {
            if (g_glow[i].m_pEntity == 0) {
                continue;
	    }
            CBaseEntity ent;
            if (!m_mem.Read(g_glow[i].m_pEntity, &ent)) {
                continue;
	    }

            if (Radar) {
                this->Radar(g_glow[i].m_pEntity, myTeam, ent.m_iTeamNum);
            }

            //switch(ent.m_iTeamNum) {
	    //    case TEAM_T || TEAM_CT:
	    //        // TODO: Uhhh shouldnt something be here?
            //        if (ent.m_iTeamNum == myTeam) {
            //            continue;
            //        } else if (ent.m_iTeamNum != myTeam) {
 	    //    	float dist = ent.m_vecOrigin.DistTo(vecEyes);
            //		if (dist < 775) {
	    //    	    Log("[rogue] entity within 375 distance.");
            //    	    // char m_iName[MAX_PATH]; //0x188
	    //    	    Log("[rogue] entity owned by: " + std::string(ent.m_iName));
            //                if (ent.m_iHealth < 1) {
            //                    continue;
	    //    	    }
	    //    	}
            //        } else {
            //            continue;
            //        }
            //        break;
            //}
            g_remote[writeCount].iov_base = ((uint8_t*)GlowData + (sizeof(glow_t) * i)) + glow_t::WriteStart();
            g_local[writeCount].iov_base = ((uint8_t*)&g_glow[i]) + glow_t::WriteStart();
            g_remote[writeCount].iov_len = glow_t::WriteSize();
            g_local[writeCount].iov_len = glow_t::WriteSize();

            writeCount++;
        }
        m_mem.WriteMulti(g_local, g_remote, writeCount);
        WaitMs(3);
    }
    Log("[rogue] Radar hack stopped");
}
