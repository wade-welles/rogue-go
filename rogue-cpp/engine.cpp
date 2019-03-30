#include "engine.h"
#include "offsets.h"

Engine &Engine::GetInstance()
{
    static Engine instance;
    return instance;
}

Engine::~Engine() {
}

void Engine::SetProcessManager(Process *proc)
{
    m_proc = proc;
}

CBaseEntityList Engine::GetEntityList()
{
    std::lock_guard<std::mutex> guard(m_entitymutex);
    return this->m_entitylist;
}

bool Engine::GetEntityById(int id, CBaseEntity* ent)
{
    std::lock_guard<std::mutex> guard(m_entitymutex);
    uintptr_t aEntity = m_entitylist.GetEntityPtrById(id);
    if (aEntity != 0) {
        if (m_proc->Read(aEntity, ent)) {
            return true;
        }
    }
    return false;
}

bool Engine::GetEntityPtrById(int id, uintptr_t* out)
{
    std::lock_guard<std::mutex> guard(m_entitymutex);
    *out = m_entitylist.GetEntityPtrById(id);
    return (*out != 0);
}

void Engine::UpdateEntityList()
{
    std::lock_guard<std::mutex> guard(m_entitymutex);
    m_entitylist.Reset();
    CEntInfo info;
    m_proc->Read(Offset::Client::EntityList, &info);
    if (info.m_pPrev != NULL) {
        return;
    }
    while (info.m_pNext != NULL && IsConnected()) {
        CBaseEntity ent;
        if (!m_proc->Read(info.m_pEntity, &ent)) {
            break;
        }
        m_entitylist.AddEntInfo(ent.index, info);
        if (!m_proc->Read(info.m_pNext, &info)) {
            break;
        }
    }
}

bool Engine::IsConnected()
{
    int sos = 0;
    m_proc->Read(Offset::Engine::ClientState + Offset::Static::SignOnState, &sos);
    return (sos == 6);
}

void Engine::Update()
{
    if (m_proc == nullptr) {
        return;
    }
    m_updateTick++;

    if (m_updateTick >= 20) {
        m_updateTick = 0;
        UpdateEntityList();
    }
}
