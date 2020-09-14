<template>
  <div>
    <h2>Client specific information</h2>
    <v-text-field clearable placeholder="Enter a client ID"></v-text-field>

    <client />
  </div>
</template>

<script>
import Client from "@/components/Client";

export default {
  components: {
    Client
  }
};
</script>
