#pragma once

#include "process.h"
#include <cstdint>

class Signatures {
    public:
        static bool Find(Process &mem);
        static void Print(Process &mem);
};

namespace Offset {
    namespace Client {
        extern uintptr_t EntityList;
        extern uintptr_t LocalPlayer;
        extern uintptr_t GlowObjectManager;
        extern uintptr_t PlayerResources;
	extern uintptr_t PostProcessing;
    };
    namespace Engine {
        extern uintptr_t ClientState;
    };

    namespace Static {
        constexpr size_t SignOnState = 0x1a0;
	constexpr size_t ViewAngles = 0x8E98;
	constexpr size_t BoneMatrix = 0x2c54 + 0x2c;
        constexpr size_t BoneDistance = 0x30; // Read(BoneMatrix) + BoneDistance
    };
};

namespace Netvar {
	namespace CBaseEntity {
		constexpr size_t index = 0x94;
		constexpr size_t m_vecOrigin = 0x170;
		constexpr size_t m_fFlags = 0x13c;
		constexpr size_t m_iTeamNum = 0x12c;
		constexpr size_t m_vecViewOffset = 0x140;
		constexpr size_t m_angRotation = 0x164;
		constexpr size_t m_nModelIndex = 0x290;
		constexpr size_t m_lifeState = 0x297;
	};
	namespace CBasePlayer {
		constexpr size_t m_bSpotted = 0xecd;
		constexpr size_t m_Local = 0x3700;
		constexpr size_t m_flFlashDuration = 0xad00;
		constexpr size_t m_iCrosshairID = 0xbbe4;
		namespace Local {
			constexpr size_t m_aimPunchAngle = m_Local + 0x74;
		};
	};
};
