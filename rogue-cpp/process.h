#pragma once

#include <string>
#include <vector>
#include <unordered_map>
#include <sys/uio.h>

class Process {
    public:
        Process(const std::string &processName);
        ~Process() = default;
        bool Attach();
        bool IsValid();
        bool ParseModules();
        uintptr_t GetCallAddress(uintptr_t address);
        uintptr_t GetAbsoluteAddress(uintptr_t address, size_t offset, size_t size);
        uintptr_t GetModuleStart(const std::string &moduleName);
        uintptr_t FindInModule(const std::string &moduleName, const std::string &pattern, size_t offset);
        bool HasModule(const std::string &moduleName);

        template <class T>
        inline bool Read(void* addr, T* out = NULL, size_t len = sizeof(T)) {
            struct iovec local = {out, len};
            struct iovec remote = {addr, len};
            return (process_vm_readv(m_pid, &local, 1, &remote, 1, 0) ==
                    static_cast<ssize_t>(len));
        }

        template <class T>
        inline bool Read(uintptr_t addr, T* out = NULL, size_t len = sizeof(T))
        {
            return this->Read(reinterpret_cast<void*>(addr), out, len);
        }

        template <class T>
        inline bool Write(uintptr_t addr, T value, size_t len = sizeof(T))
        {
            struct iovec local = {reinterpret_cast<void*>(&value), len};
            struct iovec remote = {reinterpret_cast<void*>(addr), len};
            return (process_vm_writev(m_pid, &local, 1, &remote, 1, 0) ==
                    static_cast<ssize_t>(len));
        }

        inline void WriteMulti(struct iovec *local, struct iovec *remote, size_t count)
        {
            process_vm_writev(m_pid, local, count, remote, count, 0);
        }

    private:
        struct Module {
            uintptr_t start;
            uintptr_t end;
        };
        std::unordered_map<std::string, Module> m_modules;
        long m_pid = -1;
        char m_lookup[255];
        std::string m_processName;
        std::vector<uint8_t> HexToBin(const std::string &str);
        void InitHexLookup();
};

