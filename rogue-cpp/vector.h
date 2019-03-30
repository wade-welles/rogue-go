#pragma once

#include <cmath>
#include <cstdint>

template <typename T>
class CUtlVector {
    public:
        size_t Count() { return m_nMaxSize; }
        uintptr_t Data() { return m_pData; }
    private:
        uintptr_t m_pData;
        uint32_t m_nMaxSize;
        uint32_t unk1;
        uint32_t m_nCount;
        uint32_t m_pDataPtrBack;
};

class Vector {
    public:
        float x, y, z;
        Vector(void);
        Vector(float ix, float iy, float iz);
        Vector(const Vector &vOther);
        Vector(const float* pFloat);
        
        void Init(float ix = 0.f, float iy = 0.f, float iz = 0.f);

        bool IsValid() const;

        float operator[](int i) const;
        float& operator[](int i);

        bool operator==(const Vector& v) const;
        bool operator!=(const Vector& v) const;

        Vector& operator+=(const Vector& v);
        Vector& operator-=(const Vector& v);
        Vector& operator*=(const Vector& v);
        Vector& operator*=(float s);
        Vector& operator/=(const Vector& v);
        Vector& operator/=(float s);

        Vector& operator=(const Vector& vOther);

        Vector operator+(const Vector& v) const;
        Vector operator-(const Vector& v) const;

        float DistTo(const Vector& vOther) const;
        float DistToSqr(const Vector& vOther) const;
        float Length(void) const;
        float LengthSqr(void) const;
        Vector Normalize(void) const;
        Vector& NormalizeInPlace(void);
};

inline void VectorAdd(const Vector& a, const Vector& b, Vector& c)
{
    c.x = a.x + b.x;
    c.y = a.y + b.y;
    c.z = a.z + b.z;
}

inline void VectorSubtract(const Vector& a, const Vector& b, Vector& c)
{
    c.x = a.x - b.x;
    c.y = a.y - b.y;
    c.z = a.z - b.z;
}

inline void VectorNormalize(Vector &v)
{
    float len = v.Length();
    if (len != 0.0f) {
        v.x /= len;
        v.y /= len;
        v.z /= len;
    } else {
        v.x /= len;
        v.y = 1.0f;
        v.z /= len;
    }
}

inline float VectorLength(const Vector& v)
{
    return (float)std::sqrt(v.x*v.x + v.y*v.y + v.z*v.z);
}

inline Vector::Vector(void)
{
    this->x = 0.f;
    this->y = 0.f;
    this->z = 0.f;
}

inline Vector::Vector(float ix, float iy, float iz)
{
    this->x = ix;
    this->y = iy;
    this->z = iz;
}

inline Vector::Vector(const float *pFloat)
{
    this->x = pFloat[0];
    this->y = pFloat[1];
    this->z = pFloat[2];
}

inline Vector::Vector(const Vector& vOther)
{
    this->x = vOther.x;
    this->y = vOther.y;
    this->z = vOther.z;
}

inline void Vector::Init(float ix, float iy, float iz)
{
    this->x = ix;
    this->y = iy;
    this->z = iz;
}

inline Vector& Vector::operator=(const Vector& vOther)
{
    this->x = vOther.x;
    this->y = vOther.y;
    this->z = vOther.z;
    return *this;
}

inline float& Vector::operator[](int i)
{
    return ((float*)this)[i];
}

inline float Vector::operator[](int i) const
{
    return ((float*)this)[i];
}

inline bool Vector::operator==(const Vector& src) const
{
    return ((this->x == src.x) && (this->y == src.y));
}

inline bool Vector::operator!=(const Vector& src) const
{
    return ((this->x != src.x) || (this->y != src.y));
}

inline Vector& Vector::operator+=(const Vector& v)
{
    this->x += v.x;
    this->y += v.y;
    this->z += v.z;
    return *this;
}

inline Vector& Vector::operator-=(const Vector& v)
{
    this->x -= v.x;
    this->y -= v.y;
    this->z -= v.z;
    return *this;
}

inline Vector& Vector::operator*=(float s)
{
    this->x *= s;
    this->y *= s;
    this->z *= s;
    return *this;
}

inline Vector& Vector::operator*=(const Vector& v)
{
    this->x *= v.x;
    this->y *= v.y;
    this->z *= v.z;
    return *this;
}

inline Vector& Vector::operator/=(float s)
{
    float oofl = 1.0f / s;
    this->x *= oofl;
    this->y *= oofl;
    this->z *= oofl;
    return *this;
}

inline Vector& Vector::operator/=(const Vector& v)
{
    this->x /= v.x;
    this->y /= v.y;
    this->z /= v.z;
    return *this;
}

inline Vector Vector::operator+(const Vector& v) const
{
    Vector res;
    VectorAdd(*this, v, res);
    return res;
}

inline Vector Vector::operator-(const Vector& v) const
{
    Vector res;
    VectorSubtract(*this, v, res);
    return res;
}

inline float Vector::DistTo(const Vector& vOther) const
{
    Vector delta;
    VectorSubtract(*this, vOther, delta);
    return delta.Length();
}

inline float Vector::DistToSqr(const Vector& vOther) const
{
    Vector delta;
    VectorSubtract(*this, vOther, delta);
    return delta.LengthSqr();
}

inline float Vector::LengthSqr(void) const
{
    return (this->x*this->x + this->y*this->y + this->z*this->z);
}

inline float Vector::Length(void) const
{
    return VectorLength(*this);
}

inline Vector Vector::Normalize(void) const
{
    Vector res = *this;
    VectorNormalize(res);
    return res;
}

inline Vector& Vector::NormalizeInPlace(void)
{
    VectorNormalize(*this);
    return *this;
}
