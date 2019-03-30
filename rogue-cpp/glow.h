#pragma once

#include "process.h"
#include "entity_list.h"
#include "glow_base.h"
#include "entity_list.h"

#include <cstddef>
#include <cstring>

#define END_OF_FREE_LIST -1
#define ENTRY_IN_USE -2

struct GlowObjectDefinition_t
{
    static inline size_t WriteStart() {
        return offsetof(struct GlowObjectDefinition_t, m_flGlowColor);
    }

    static inline size_t WriteSize() {
        return offsetof(struct GlowObjectDefinition_t, m_nSplitScreenSlot) -
               offsetof(struct GlowObjectDefinition_t, m_flGlowColor);
    }

    void SetColor(const float rgba[4]) {
        constexpr int clrSize = sizeof(float) * 4;
        memcpy(m_flGlowColor, rgba, clrSize);
    }

    void SetRender(bool occluded, bool unoccluded)
    {
        m_bRenderWhenOccluded = occluded;
        m_bRenderWhenUnoccluded = unoccluded;
        m_nSplitScreenSlot = -1;
        m_flBloomAmount = 1.f;
    }

    uintptr_t m_pEntity;
    float m_flGlowColor[3];
    float m_flGlowAlpha;
    char  unknown[4];
    float m_flMovementAlpha;
    float m_flBloomAmount;
    float m_localPlayerIsZeroPoint3;
    bool  m_bRenderWhenOccluded;
    bool  m_bRenderWhenUnoccluded;
    bool  m_bFullBloomRender;
    char  m_unknown[1];
    int   m_nFullBloomStencilTestValue;
    int   m_GlowStyle;
    int   m_nSplitScreenSlot;
    int   m_nNextFreeSlot;
    }; // sizeof() == 0x3c

class CGlowObjectManager
{
public:
    size_t Count() {
        return m_GlowObjectDefinitions.Count();
    }

    uintptr_t Data() {
        return m_GlowObjectDefinitions.Data();
    }
    unsigned int UNK() {
        return m_GlowObjectDefinitions.Count();
    }
private:
    CUtlVector<GlowObjectDefinition_t> m_GlowObjectDefinitions;
    int m_nFirstFreeSlot;
};

class Glow final : public GlowBase {
    Process &m_mem;
    public:
        Glow(Process &mem) : m_mem(mem) {};
    protected:
        void Run();
    private:
        void Radar(uintptr_t entPtr, int lTeam, int rTeam);
};
