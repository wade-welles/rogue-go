#include "glow_base.h"

#include <chrono>
#include <iostream>

GlowBase::GlowBase()
{
    m_stop = false;
}

GlowBase::~GlowBase()
{
    Stop();
}

void GlowBase::Start()
{
    Stop();
    m_stop = false;
    m_thread = std::thread(&GlowBase::Run, this);
}

void GlowBase::Stop()
{
    m_stop = true;
    if (m_thread.joinable()) {
        m_thread.join();
    }
}

void GlowBase::Log(const std::string &msg)
{
    std::cout << msg << std::endl;
}

void GlowBase::LogWait(const std::string &msg, size_t timeout)
{
    Log(msg);
    Wait(timeout);
}
