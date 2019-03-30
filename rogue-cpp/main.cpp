#include "glow.h"
#include "entity_list.h"
#include "types.h"
#include "engine.h"
#include "offsets.h"
#include "process.h"

#include <chrono>
#include <cstring>
#include <cstdlib>
#include <iostream>
#include <thread>

#include <signal.h>
#include <unistd.h>

#define LOG(X) std::cout << X << std::flush


bool shouldQuit = false;

void exitHandle(int)
{
    shouldQuit = true;
}

void connectSignals(struct sigaction &handle)
{
    handle.sa_handler = exitHandle;
    sigemptyset(&handle.sa_mask);
    handle.sa_flags = 0;
    sigaction(SIGINT, &handle, NULL);
    sigaction(SIGQUIT, &handle, NULL);
}

int main()
{
    if (getuid() != 0) {
        LOG("This program must be ran as root.\n");
        return 0;
    }

    char* displayName = getenv("DISPLAY");
    if (displayName) {
        printf("Display is: %s\n", displayName);
    } else {
        printf("Failed to find display!\n");
    }

    struct sigaction ccHandle;
    connectSignals(ccHandle);
    
    Process proc(PROCESS_NAME);
    
    LOG("Waiting for process...");
    
    while (!proc.Attach() && !shouldQuit) {
        std::this_thread::sleep_for(std::chrono::seconds(1));
    }

    LOG("Done.\nWaiting for client and engine library...");

    while (!shouldQuit) {
        proc.ParseModules();
        if (proc.HasModule(CLIENT_SO) && proc.HasModule(ENGINE_SO)) {
            break;
        }
        std::this_thread::sleep_for(std::chrono::seconds(1));
    }

    if (shouldQuit) {
        return 0;
    }

    LOG("Done.\n");
    Signatures::Find(proc);
    Signatures::Print(proc);

    auto& eng = Engine::GetInstance();
    eng.SetProcessManager(&proc);
    eng.Update();

    // Feature handlers
    ///////////////////////////////////////////////////////////////////////////
    Glow glow(proc);

    while (!shouldQuit) {
        if (!proc.IsValid()) {
            shouldQuit = true;
            LOG("Lost connection to process... Exiting.\n");
            break;
        }

        // ### BEGIN IN-GAME HACKS ###
        if (eng.IsConnected()) {
            glow.Start();

            while (eng.IsConnected() && !shouldQuit) {
                eng.Update();
                std::this_thread::sleep_for(std::chrono::milliseconds(50));
            }

            glow.Stop();
        }
        // ### END IN-GAME HACKS ###
        std::this_thread::sleep_for(std::chrono::seconds(1));
    }
    return 0;
}
