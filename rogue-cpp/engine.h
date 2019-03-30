#pragma once

#include "entity_list.h"
#include "glow.h"
#include "process.h"

#include <iostream>
#include <mutex>
#include <thread>

class Engine final {
    public:
        static Engine &GetInstance();
        void Update();
        void SetProcessManager(Process *proc);
        bool IsConnected();

        // Entity List
        CBaseEntityList GetEntityList();
        bool GetEntityById(int id, CBaseEntity* ent);
        bool GetEntityPtrById(int id, uintptr_t* out);

        // Useful things
    private:
        Engine() = default;
        ~Engine();
        Engine(const Engine&) = delete;
        Engine& operator=(const Engine&) = delete;
        Engine(Engine&&) = delete;
        Engine& operator=(Engine&&) = delete;
        void UpdateEntityList();

        Process *m_proc = nullptr;
        CBaseEntityList m_entitylist;
        size_t m_updateTick = 0;
        std::mutex m_entitymutex;
};
