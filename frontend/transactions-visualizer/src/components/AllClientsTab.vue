<template>
  <div>
    <h2>All clients information</h2>
    <v-simple-table>
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">Name</th>
            <th class="text-left">Age</th>
            <th class="text-left">ID</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="item in buyers" :key="item.name">
            <td>{{ item.name }}</td>
            <td>{{ item.age }}</td>
            <td>{{ item.id }}</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>

    <div class="text-center">
      <v-btn
        class="ma-5"
        :dark="hasMore"
        :loading="loading"
        :disabled="!hasMore"
        @click="loadBuyers"
        >Load more</v-btn
      >
    </div>
  </div>
</template>

<script>
import Endpoints from "../services/EndpointsService";
import EventBus from "../event-bus";

const PAGE_LIMIT = 12;

export default {
  data: () => ({
    buyers: [],
    hasMore: true,
    page: 1,
    loading: false
  }),

  methods: {
    async loadBuyers() {
      this.loading = true;
      const { data, success } = await Endpoints.getBuyers(
        this.page,
        PAGE_LIMIT
      );
      this.loading = false;

      if (success) {
        this.buyers = this.buyers.concat(data);
        this.page = this.page + 1;
        this.hasMore = data.length >= PAGE_LIMIT;
      } else {
        EventBus.$emit("alert", {
          alert: true,
          type: "error",
          message: data
        });
      }
    }
  },

  mounted() {
    this.loadBuyers();
  }
};
</script>
