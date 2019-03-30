package rogue

type MemorySignature struct {
	Hex    string
	Offset uint
}

// TODO: Use offsets.go to automatically set the offsets for each of these. There is a great C++ example to reference for help. 

//SIGNATURE(LocalPlayer, "48 89 E5 74 0E 48 8D 05", 0x7);
var localPlayer = MemorySignature{
	Hex:    "48 89 E5 84 0E 48 8D 05",
	Offset: 0x7,
}

//SIGNATURE(GlowObjectManager, "E8 ?? ?? ?? ?? 48 8B 3D ?? ?? ?? ?? BE 01 00 00 00 C7", 0x0);
var glowManager = MemorySignature{
	Hex:    "E8 ?? ?? ?? ?? 48 8B 3D ?? ?? ?? ?? BE 01 00 00 00 C7",
	Offset: 0x0,
}

//SIGNATURE(PlayerResources, "48 8B 05 ?? ?? ?? ?? 55 48 89 E5 48 85 C0 74 10 48", 0x2);
var resources = MemorySignature{
	Hex:    "48 8B 05 ?? ?? ?? ?? 55 48 89 E5 48 85 C0 74 10 48",
	Offset: 0x2,
}

//SIGNATURE(EntityList, "55 48 89 E5 48 83 EC 10 8B 47 34 48 8D 75 F0 89 45 F0 48 8B 05 ?? ?? ?? ?? 48 8B 38", 0x12);
var entities = MemorySignature{
	Hex:    "55 48 89 E5 48 83 EC 10 8B 47 34 48 8D 75 F0 48 8B 05 ?? ?? ?? ?? 48 8B 38",
	Offset: 0x12,
}

//SIGNATURE(ClientState, "48 8B 05 ?? ?? ?? ?? 55 48 8D 3D ?? ?? ?? ?? 48 89 E5 FF 50 28", 0x0);
var state = MemorySignature{
	Hex:    "48 8B 05 ?? ?? ?? ?? 55 48 8D 3D ?? ?? ?? ?? 48 89 E5 FF 50 28",
	Offset: 0x0,
}
