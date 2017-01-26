#include <cps.h>

char** MakeCharArray(int size) {
        return calloc(sizeof(char*), size);
}

void SetArrayString(char **a, char *s, int n) {
        a[n] = s;
}

void FreeCharArray(char **a, int size) {
        int i;
        for (i = 0; i < size; i++) {
                free(a[i]);
	}
        free(a);
}


