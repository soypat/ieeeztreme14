var _, o = [],
    c = 0,
    p = 0,
    j = 0,
    i = [88, 88, 88],
    m = new Uint8Array(30000);

function q(i) {
    self.postMessage({
        o: [i]
    })
}
i.length && (m[p + 1] = i.pop());
i.length && (m[p + 2] = i.pop());
i.length && (m[p + 3] = i.pop());
if ((_ = m[p + 2]) !== 0) {
    m[p + 3] += _ * -1;
    m[p + 4] += _;
    m[p + 2] = 0;
}
m[p + 2]++;
m[p + 2]--;
if ((_ = m[p + 4]) !== 0) {
    m[p + 2] += _;
    m[p + 4] = 0;
}
m[p + 3]++;
m[p + 3]--;
if ((_ = m[p + 1]) !== 0) {
    m[p + 2] += _ * -1;
    m[p] += _;
    m[p + 1] = 0;
}
if ((_ = m[p]) !== 0) {
    m[p + 1] += _;
    m[p] = 0;
}
m[p + 6] += 2;
m[p + 6]--;
m[p + 8]++;
m[p + 8]--;
m[p + 13]++;
while (m[p + 13] !== 0) {
    if ((_ = m[p + 17]) !== 0) {
        m[p + 17] = 0;
    }
    if ((_ = m[p + 16]) !== 0) {
        m[p + 16] = 0;
    }
    if ((_ = m[p + 15]) !== 0) {
        m[p + 15] = 0;
    }
    m[p + 15]--;
    m[p + 15] += 2;
    m[p + 15]--;
    m[p + 15]++;
    m[p + 15]--;
    if ((_ = m[p + 14]) !== 0) {
        m[p + 14] = 0;
    }
    if ((_ = m[p + 13]) !== 0) {
        m[p + 13] = 0;
    }
    m[p + 9]++;
    m[p + 9]--;
    m[p + 6]--;
    m[p + 6]++;
    m[p + 3]++;
    m[p + 3]--;
    if ((_ = m[p + 1]) !== 0) {
        m[p + 8] += _;
        m[p + 14] += _;
        m[p + 1] = 0;
    }
    if ((_ = m[p + 2]) !== 0) {
        m[p + 9] += _;
        m[p + 15] += _;
        m[p + 2] = 0;
    }
    if ((_ = m[p + 3]) !== 0) {
        m[p + 10] += _;
        m[p + 16] += _;
        m[p + 3] = 0;
    }
    m[p + 5]--;
    m[p + 5] += 2;
    m[p + 5]--;
    m[p + 5]++;
    m[p + 6]++;
    m[p + 6]--;
    m[p + 6]++;
    while (m[p + 6] !== 0) {
        m[p + 5]--;
        p -= 1;
    }
    while (m[p + 5] !== 0) {
        m[p + 5]--;
        m[p + 7] += 2;
        m[p + 7]--;
        p -= 1;
    }
    if ((_ = m[p + 9]) !== 0) {
        m[p + 2] += _;
        m[p + 9] = 0;
    }
    if ((_ = m[p + 10]) !== 0) {
        m[p + 3] += _;
        m[p + 10] = 0;
    }
    if ((_ = m[p + 11]) !== 0) {
        m[p + 4] += _;
        m[p + 11] = 0;
    }
    m[p + 18]++;
    while (m[p + 18] !== 0) {
        m[p + 18]--;
        m[p + 14]++;
        m[p + 14] -= 2;
        m[p + 14]++;
        m[p + 13]++;
        m[p + 13]--;
        if ((_ = m[p + 7]) !== 0) {
            m[p + 9] += _;
            m[p + 11] += _;
            m[p + 7] = 0;
        }
        m[p + 7]++;
        m[p + 7]--;
        if ((_ = m[p + 8]) !== 0) {
            m[p + 10] += _;
            m[p + 12] += _;
            m[p + 8] = 0;
        }
        if ((_ = m[p + 9]) !== 0) {
            m[p + 7] += _;
            m[p + 9] = 0;
        }
        if ((_ = m[p + 10]) !== 0) {
            m[p + 8] += _;
            m[p + 10] = 0;
        }
        m[p + 10]++;
        while (m[p + 10] !== 0) {
            if ((_ = m[p + 10]) !== 0) {
                m[p + 10] = 0;
            }
            m[p + 10]--;
            m[p + 10] += 2;
            while (m[p + 11] !== 0) {
                m[p + 10]--;
                p -= 1;
            }
            m[p + 10]++;
            m[p + 10]--;
            while (m[p + 10] !== 0) {
                m[p + 10]++;
                m[p + 10] -= 2;
                m[p + 12]--;
                p -= 1;
            }
            m[p + 12]--;
            m[p + 12]++;
            m[p + 12]--;
            m[p + 15]++;
            while (m[p + 16] !== 0) {
                m[p + 15]--;
                p -= 1;
            }
            m[p + 15]++;
            m[p + 15]--;
            while (m[p + 15] !== 0) {
                m[p + 15]--;
                m[p + 16]++;
                while (m[p + 17] !== 0) {
                    m[p + 17]--;
                    m[p + 16]--;
                    p -= 1;
                }
                while (m[p + 16] !== 0) {
                    while (m[p + 18] !== 0) {
                        m[p + 18]--;
                        m[p + 16]--;
                        m[p + 16]++;
                        m[p + 16]--;
                        p -= 1;
                    }
                    m[p + 18]--;
                    while (m[p + 16] !== 0) {
                        m[p + 16]--;
                        m[p + 17]--;
                        m[p + 18]++;
                        m[p + 18]--;
                        m[p + 19]--;
                        m[p + 19]++;
                        m[p + 19]--;
                        p -= 1;
                    }
                }
            }
            m[p + 17]--;
            m[p + 12]++;
            while (m[p + 13] !== 0) {
                m[p + 12]--;
                p -= 1;
            }
            while (m[p + 12] !== 0) {
                m[p + 13]++;
                m[p + 13]--;
                while (m[p + 14] !== 0) {
                    m[p + 14]--;
                    m[p + 14]++;
                    m[p + 12]--;
                    p -= 1;
                }
                while (m[p + 12] !== 0) {
                    p -= 1;
                }
            }
            m[p + 13]--;
            m[p + 13]++;
            m[p + 13]--;
            p += 3;
        }
        m[p + 18]++;
    }
    m[p + 18]--;
    m[p + 16]++;
    m[p + 16]--;
    m[p + 14]++;
    m[p + 15]--;
    m[p + 15] += 2;
    while (m[p + 15] !== 0) {
        m[p + 14]--;
        p -= 1;
    }
    while (m[p + 14] !== 0) {
        m[p + 14]--;
        m[p + 14]++;
        m[p + 16]++;
        while (m[p + 16] !== 0) {
            m[p + 14]--;
            m[p + 15]--;
            m[p + 15]++;
            p -= 1;
        }
        while (m[p + 14] !== 0) {
            m[p + 17]++;
            while (m[p + 17] !== 0) {
                m[p + 16]--;
                m[p + 16]++;
                m[p + 14]--;
                p -= 1;
            }
            while (m[p + 14] !== 0) {
                p -= 1;
            }
        }
    }
    m[p + 15]--;
    p += 2;
}
if ((_ = m[p + 17]) !== 0) {
    m[p + 17] = 0;
}
m[p + 10]--;
m[p + 10]++;
if ((_ = m[p + 3]) !== 0) {
    m[p + 3] = 0;
}
m[p + 5]++;
m[p + 6]++;
while (m[p + 6] !== 0) {
    m[p + 5]--;
    p -= 1;
}
while (m[p + 5] !== 0) {
    m[p + 5]--;
    m[p + 7]++;
    p -= 1;
}
m[p + 6]++;
while (m[p + 7] !== 0) {
    m[p + 6]--;
    p -= 1;
}
while (m[p + 6] !== 0) {
    while (m[p + 8] !== 0) {
        m[p + 6]--;
        p -= 1;
    }
    while (m[p + 6] !== 0) {
        p -= 1;
    }
}
m[p + 7]--;
while (m[p + 7] !== 0) {
    m[p + 7]++;
    m[p + 13] += 2;
    m[p + 13]--;
    m[p + 13] += 9;
    m[p + 7]++;
    m[p + 8]++;
    m[p + 8]--;
    while (m[p + 8] !== 0) {
        m[p + 7]--;
        p -= 1;
    }
    while (m[p + 7] !== 0) {
        while (m[p + 9] !== 0) {
            m[p + 7]--;
            p -= 1;
        }
        while (m[p + 7] !== 0) {
            p -= 1;
        }
    }
    m[p + 8]--;
    while (m[p + 8] !== 0) {
        m[p + 8] += 2;
        while (m[p + 9] !== 0) {
            m[p + 8]--;
            p -= 1;
        }
        while (m[p + 8] !== 0) {
            m[p + 8]--;
            m[p + 10]--;
            p -= 1;
        }
        m[p + 10]--;
        m[p + 12]++;
        m[p + 14]++;
        m[p + 15]--;
        while (m[p + 15] !== 0) {
            m[p + 14]--;
            p -= 1;
        }
        while (m[p + 14] !== 0) {
            if ((_ = m[p + 12]) !== 0) {
                m[p + 15] += _;
                m[p + 12] = 0;
            }
            m[p + 12]++;
            m[p + 12]--;
            m[p + 17]++;
            m[p + 18]++;
            while (m[p + 18] !== 0) {
                m[p + 17]--;
                p -= 1;
            }
            while (m[p + 17] !== 0) {
                m[p + 17]--;
                m[p + 19]++;
                p -= 1;
            }
            m[p + 15]--;
        }
        m[p + 10]++;
        while (m[p + 11] !== 0) {
            m[p + 10]--;
            p -= 1;
        }
        while (m[p + 10] !== 0) {
            while (m[p + 12] !== 0) {
                m[p + 10]--;
                p -= 1;
            }
            while (m[p + 10] !== 0) {
                p -= 1;
            }
        }
        m[p + 11]--;
        p += 3;
    }
    m[p + 10] += 8;
    if ((_ = m[p + 10]) !== 0) {
        m[p + 11] += _ * 6;
        m[p + 10] = 0;
    }
    if ((_ = m[p + 11]) !== 0) {
        m[p + 7] += _;
        m[p + 11] = 0;
    }
    m[p + 16]++;
    m[p + 17]--;
    m[p + 17]++;
    while (m[p + 17] !== 0) {
        m[p + 16]--;
        p -= 1;
    }
    while (m[p + 16] !== 0) {
        while (m[p + 18] !== 0) {
            m[p + 16]--;
            p -= 1;
        }
        while (m[p + 16] !== 0) {
            p -= 1;
        }
    }
    m[p + 17]--;
    while (m[p + 17] !== 0) {
        m[p + 17] += 2;
        while (m[p + 18] !== 0) {
            m[p + 17]--;
            p -= 1;
        }
        while (m[p + 17] !== 0) {
            m[p + 17]--;
            m[p + 19]--;
            p -= 1;
        }
        m[p + 19]--;
        m[p + 11]++;
        m[p + 12]++;
        while (m[p + 12] !== 0) {
            m[p + 11]--;
            p -= 1;
        }
        while (m[p + 11] !== 0) {
            m[p + 11]--;
            m[p + 13]++;
            p -= 1;
        }
        m[p + 19]++;
        while (m[p + 20] !== 0) {
            m[p + 19]--;
            p -= 1;
        }
        while (m[p + 19] !== 0) {
            while (m[p + 21] !== 0) {
                m[p + 19]--;
                p -= 1;
            }
            while (m[p + 19] !== 0) {
                p -= 1;
            }
        }
        m[p + 20]--;
        p += 3;
    }
    if ((_ = m[p + 15]) !== 0) {
        m[p + 15] = 0;
    }
    m[p + 10]++;
    while (m[p + 11] !== 0) {
        m[p + 10]--;
        p -= 1;
    }
    while (m[p + 10] !== 0) {
        while (m[p + 12] !== 0) {
            m[p + 10]--;
            p -= 1;
        }
        while (m[p + 10] !== 0) {
            p -= 1;
        }
    }
    m[p + 11]--;
    p += 4;
}
while (m[p + 5] !== 0) {
    q(m[p + 5]);
    p -= 1;
}

return self.postMessage({
    s: -1,
    o: o,
    c: c,
    m: m,
    p: p + 5,
    n: -1
});