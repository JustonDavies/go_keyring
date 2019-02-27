//-- Includes -----------------------------------------------------------------------------------------------------------
#include "linux_secret_service.h"

//-- Structs -----------------------------------------------------------------------------------------------------------

//-- Exported Functions ------------------------------------------------------------------------------------------------
GError *getSecretByAttribute(char *name, char *value, char **secret) {
	//-- Declarations ----------
	GError *err;

	GHashTable *attributes;
	GList *results;
	SecretService *service;

  //-- Initializing ----------
  {
    err = NULL;
	  service = secret_service_get_sync(SECRET_SERVICE_NONE, NULL, &err);

  	attributes = g_hash_table_new_full(g_str_hash, g_str_equal, g_free, g_free);
    g_hash_table_insert(attributes, g_strdup(name), g_strdup(value));
  }

  //-- Query ----------
	if (err == NULL) {
		results = secret_service_search_sync(
		  service,
		  NULL,
		  attributes,
		  SECRET_SEARCH_LOAD_SECRETS | SECRET_SEARCH_ALL | SECRET_SEARCH_UNLOCK,
		  NULL,
		  &err
		);
  }

  //-- Parse/Populate ----------
	if (err == NULL) {
	  for (GList *i = results; i != NULL; i = g_list_next(i)) {
	    SecretValue *temp = secret_item_get_secret(i->data);
			if (secret != NULL) { *secret = strdup(secret_value_get(temp, NULL)); }
			secret_value_unref(temp);
	  }
	}

	//-- Frees ----------
	{
	  g_hash_table_unref(attributes);
    g_list_free_full(results, g_object_unref);
    g_object_unref(service);
  }

  //-- Return ----------
	return err;
}

//-- Internal Functions ------------------------------------------------------------------------------------------------
