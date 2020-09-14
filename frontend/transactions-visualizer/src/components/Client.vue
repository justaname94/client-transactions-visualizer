<template>
  <div>
    <h2>{{ client.buyer.name }} {{ client.buyer.age }}</h2>
    <v-divider class="my-4"></v-divider>
    <h2 class="mb-3">Transactions</h2>

    <div v-for="transaction in client.buyer.transaction" :key="transaction.id">
      <div class="text-left">
        <v-chip class="ma-2 px-4" color="primary">
          <v-icon left>mdi-identifier</v-icon>{{ transaction.id }}
        </v-chip>

        <v-chip class="ma-2 px-4" color="secondary">
          <v-icon left>mdi-cellphone-link</v-icon>{{ transaction.device }}
        </v-chip>

        <v-chip class="ma-2 px-4" color="red" text-color="white">
          <v-icon left>mdi-ip-network</v-icon>{{ transaction.ip }}
        </v-chip>

        <v-chip class="ma-2 px-4" color="green" text-color="white">
          <v-icon left>mdi-currency-usd</v-icon
          >{{ toDollars(transaction.total) }}
        </v-chip>
      </div>

      <v-simple-table class="table">
        <template v-slot:default>
          <thead>
            <tr>
              <th class="text-left">Name</th>
              <th class="text-left">Price ($ USD)</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="p in transaction.product" :key="p.name">
              <td>{{ p.name }}</td>
              <td>{{ toDollars(p.price) }}</td>
            </tr>
          </tbody>
        </template>
      </v-simple-table>
    </div>
    <v-divider class="my-4"></v-divider>
    <h2 class="mb-3">Related IP Buyers</h2>
    <v-simple-table>
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">Name</th>
            <th class="text-left">ID</th>
            <th class="text-left">IP</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="buyer in client.relatedIpBuyers" :key="buyer.id">
            <td>{{ buyer.name }}</td>
            <td>{{ buyer.id }}</td>
            <td>{{ buyer.ip }}</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
    <v-divider class="my-4"></v-divider>
    <h2 class="mb-3">Product recomendations</h2>
    <v-simple-table class="table">
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">Name</th>
            <th class="text-left">Price ($ USD)</th>
          </tr>
        </thead>
        <tbody>
          <tr
            v-for="recomendation in client.recommendations"
            :key="recomendation.name"
          >
            <td>{{ recomendation.name }}</td>
            <td>{{ toDollars(recomendation.price) }}</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
  </div>
</template>

<script>
export default {
  props: {
    client: {
      type: Object,
      required: true
    }
  },

  methods: {
    toDollars(cents) {
      return parseFloat(cents * 0.01).toFixed(2);
    }
  }
};
</script>

<style>
.table table {
  table-layout: fixed;
}
</style>
