#include "entity_list.h"
#include <iostream>

size_t CBaseEntityList::Count()
{
    return m_entinfo.size();
}

void CBaseEntityList::Reset()
{
    m_entinfo.clear();
}

void CBaseEntityList::AddEntInfo(int index, CEntInfo info)
{
    m_entinfo.emplace(std::make_pair(index, info));
}

uintptr_t CBaseEntityList::GetEntityPtrById(int index)
{
    try {
        CEntInfo& info = m_entinfo.at(index);
        return info.m_pEntity;
    } catch (std::out_of_range& e) {
        return 0;
    }
}
