#pragma once

struct matrix3x4_t {
    matrix3x4_t() {}
    float *operator[](int i) { return m_flMatVal[i]; }
    float m_flMatVal[3][4];
};
