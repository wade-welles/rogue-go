#pragma once

#include <cmath>

class Vector2D {
    public:
        float x, y;
        Vector2D(void);
        Vector2D(float ix, float iy);
        Vector2D(const Vector2D &vOther);
        Vector2D(const float* pFloat);
        
        void Init(float ix = 0.f, float iy = 0.f);

        bool IsValid() const;

        float operator[](int i) const;
        float& operator[](int i);

        bool operator==(const Vector2D& v) const;
        bool operator!=(const Vector2D& v) const;

        Vector2D& operator+=(const Vector2D& v);
        Vector2D& operator-=(const Vector2D& v);
        Vector2D& operator*=(const Vector2D& v);
        Vector2D& operator*=(float s);
        Vector2D& operator/=(const Vector2D& v);
        Vector2D& operator/=(float s);

        Vector2D& operator=(const Vector2D& vOther);

        float DistTo(const Vector2D& vOther) const;
        float DistToSqr(const Vector2D& vOther) const;
        float Length(void) const;
        float LengthSqr(void) const;
};

inline void Vector2DSubtract(const Vector2D& a, const Vector2D& b, Vector2D& c)
{
    c.x = a.x - b.x;
    c.y = a.y - b.y;
}

inline float Vector2DLength(const Vector2D& v)
{
    return (float)std::sqrt(v.x*v.x + v.y*v.y);
}

inline Vector2D::Vector2D(void)
{
}

inline Vector2D::Vector2D(float ix, float iy)
{
    this->x = ix;
    this->y = iy;
}

inline Vector2D::Vector2D(const float *pFloat)
{
    this->x = pFloat[0];
    this->y = pFloat[1];
}

inline Vector2D::Vector2D(const Vector2D& vOther)
{
    this->x = vOther.x;
    this->y = vOther.y;
}

inline void Vector2D::Init(float ix, float iy)
{
    this->x = ix;
    this->y = iy;
}

inline Vector2D& Vector2D::operator=(const Vector2D& vOther)
{
    this->x = vOther.x;
    this->y = vOther.y;
    return *this;
}

inline float& Vector2D::operator[](int i)
{
    return ((float*)this)[i];
}

inline float Vector2D::operator[](int i) const
{
    return ((float*)this)[i];
}

inline bool Vector2D::operator==(const Vector2D& src) const
{
    return ((this->x == src.x) && (this->y == src.y));
}

inline bool Vector2D::operator!=(const Vector2D& src) const
{
    return ((this->x != src.x) || (this->y != src.y));
}

inline Vector2D& Vector2D::operator+=(const Vector2D& v)
{
    this->x += v.x;
    this->y += v.y;
    return *this;
}

inline Vector2D& Vector2D::operator-=(const Vector2D& v)
{
    this->x -= v.x;
    this->y -= v.y;
    return *this;
}

inline Vector2D& Vector2D::operator*=(float s)
{
    this->x *= s;
    this->y *= s;
    return *this;
}

inline Vector2D& Vector2D::operator*=(const Vector2D& v)
{
    this->x *= v.x;
    this->y *= v.y;
    return *this;
}

inline Vector2D& Vector2D::operator/=(float s)
{
    float oofl = 1.0f / s;
    this->x *= oofl;
    this->y *= oofl;
    return *this;
}

inline Vector2D& Vector2D::operator/=(const Vector2D& v)
{
    this->x /= v.x;
    this->y /= v.y;
    return *this;
}

inline float Vector2D::DistTo(const Vector2D& vOther) const
{
    Vector2D delta;
    Vector2DSubtract(*this, vOther, delta);
    return delta.Length();
}

inline float Vector2D::DistToSqr(const Vector2D& vOther) const
{
    Vector2D delta;
    Vector2DSubtract(*this, vOther, delta);
    return delta.LengthSqr();
}

inline float Vector2D::LengthSqr(void) const
{
    return (this->x*this->x + this->y*this->y);
}

inline float Vector2D::Length(void) const
{
    return Vector2DLength(*this);
}
