<template>
  <div>
    <h2>Client specific information</h2>
    <div class="d-flex">
      <v-text-field
        class="mt-2"
        :loading="loading"
        v-model="text"
        placeholder="Enter a client ID"
        @keydown.enter="loadBuyer"
      ></v-text-field>
      <v-btn
        :loading="loading"
        :disabled="text === ''"
        class="ma-5"
        :dark="text !== ''"
        @click="loadBuyer"
        >Search</v-btn
      >
    </div>
    <client v-if="exists" :client="client" />
  </div>
</template>

<script>
import Client from "@/components/Client";
import Endpoints from "../services/EndpointsService";
import EventBus from "../event-bus";

export default {
  components: {
    Client
  },

  data: () => ({
    loading: false,
    text: "",
    exists: false,
    client: {}
  }),

  methods: {
    async loadBuyer() {
      const text = this.text ? this.text : "404";

      this.loading = true;
      const { data, success } = await Endpoints.getBuyer(text);
      this.loading = false;

      if (success) {
        this.exists = true;
        this.client = data;
      } else {
        this.client = {};
        this.exists = false;
        EventBus.$emit("alert", {
          alert: true,
          type: "error",
          message: data
        });
      }
    }
  }
};
</script>
