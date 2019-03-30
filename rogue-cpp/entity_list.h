#pragma once

#include "types.h"
#include "vector.h"

#include <cstdint>
#include <unordered_map>

#define MAX_PATH 260

#define LIFE_STATE_ALIVE 0
#define LIFE_STATE_DYING 1
#define LIFE_STATE_DEAD 2

#define TEAM_NONE 0
#define TEAM_SPEC 1
#define TEAM_T    2
#define TEAM_CT   3


class CBaseEntity {
    public:
        int GetHealth() const;
        int GetLifeState() const;
        int GetTeam();
        int GetIndex(); 

//    private:
        char __buf_0x00[0x40]; // 0x0
        float m_flLastMadeNoiseTime;  //0x40
        char __buf_0x44[0x4C]; //0x44
        bool m_bIsAutoaimTarget; //0x90
        char __buf_0x91[0x3]; //0x91
        int index; // 0x94
        unsigned short m_EntClientFlags; // 0x98
        const void *model; // 0xA0
        color32 m_clrRender; // 0xA8
        int m_cellbits; // 0xAC
        int m_cellwidth; // 0xB0
        int m_cellX; // 0xB4
        int m_cellY; // 0xB8
        int m_cellZ; // 0xBC
        Vector m_vecCellOrigin; // 0xC0
        Vector m_vecAbsVelocity; // 0xCC
        Vector m_vecAbsOrigin; // 0xD8
        Vector m_vecOrigin; // 0xE4
        Vector m_vecAngVelocity; // 0xF0
        QAngle m_angAbsRotation; // 0xFC
        QAngle m_angRotation; // 0x108
        float m_flGravity; // 0x114
        float m_flProxyRandomValue; // 0x118
        int m_iEFlags; // 0x11C
        unsigned char m_nWaterType; // 0x120
        bool m_bDormant; // 0x121
	char __buf_0x122[0x6]; // 0x122
        int m_fEffects; // 0x124
        int m_iTeamNum; // 0x128
        int m_iPendingTeamNum; // 0x12C
        int m_nNextThinkTick; // 0x130
        int m_iHealth; // 0x134
        int m_fFlags; // 0x138
        Vector m_vecViewOffset; // 0x13C
        Vector m_vecVelocity; // 0x148
        Vector m_vecBaseVelocity; // 0x154
        QAngle m_angNetworkAngles; // 0x160
        Vector m_vecNetworkOrigin; // 0x16C
        float m_flFriction; // 0x178
        int moveparent;  // 0x17C
        int m_hOwnerEntity; // 0x180
        int m_hGroundEntity; // 0x184
        char m_iName[MAX_PATH]; //0x188
        short m_nModelIndex; //0x28C
        unsigned char m_nRenderFX; //0x28E
        unsigned char m_nRenderMode; //0x28F
        unsigned char m_MoveType;  //0x290
        unsigned char m_MoveCollide; //0x291
        unsigned char m_nWaterLevel; //0x292
        char m_lifeState; //0x293
        float m_flAnimTime; //0x294
        float m_flOldAnimTime; //0x298
        float m_flSimulationTime; //0x29C
        float m_flOldSimulationTime; //0x2A0
        float m_flCreateTime; //0x2A4
        unsigned char m_nOldRenderMode; //0x2A8
        unsigned short m_hRender; //0x2A9
        unsigned char m_VisibilityBits; //0x2AB
        bool m_bReadyToDraw; //0x2AC
        bool m_bClientSideRagdoll; //0x2AD
        int m_nLastThinkTick; //0x2AE
        char m_takedamage; //0x2B2
        float m_flSpeed; //0x2B3
};

typedef struct {
    CBaseEntity entity;
    void* entityPtr;
} EntityInfo;

typedef struct {
    Vector head;
    Vector origin;
    Vector eyes;
    Vector aimLine;
    EntityInfo entityInfo;
    char name[256];
    int health;
    bool scoped;
    bool myTeam;
    bool defusing;
    bool reloading;
    float lby;
    bool spectatingMe;
} EntityToScreen;

class Entity {
    char data[0x1000];
};

class CEntInfo {
    public:
        uintptr_t m_pEntity;
        long m_SerialNumber;
        CEntInfo *m_pPrev;
        CEntInfo *m_pNext;
};

class CBaseEntityList {
    public:
        const std::unordered_map<int, CEntInfo> &Data() const { return m_entinfo; }
        void AddEntInfo(int index, CEntInfo einfo);
        uintptr_t GetEntityPtrById(int index);
        void Reset();
        size_t Count();
    private:
        std::unordered_map<int, CEntInfo> m_entinfo;
};
