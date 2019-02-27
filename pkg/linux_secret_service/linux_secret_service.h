#include <stdlib.h>
#include <string.h>

#include <libsecret/secret.h>

GError *getSecretByAttribute(char *name, char *value, char **secret) ;