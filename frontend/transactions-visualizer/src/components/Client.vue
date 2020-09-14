<template>
  <div>
    <h2>{{name}} {{age}}</h2>
    <v-divider class="my-4"></v-divider>
    <h2 class="mb-3">Transactions</h2>

    <div class="text-left">
      <v-chip class="ma-2 px-4" color="primary">
        <v-icon left>mdi-identifier</v-icon>00005f4d9336
      </v-chip>

      <v-chip class="ma-2 px-4" color="secondary">
        <v-icon left>mdi-cellphone-link</v-icon>Android
      </v-chip>

      <v-chip class="ma-2 px-4" color="red" text-color="white">
        <v-icon left>mdi-ip-network</v-icon>99.65.7.60
      </v-chip>

      <v-chip class="ma-2" color="green" text-color="white">
        <v-icon left>mdi-currency-usd</v-icon>248.5
      </v-chip>
    </div>

    <v-simple-table>
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">Name</th>
            <th class="text-left">Price ($ USD)</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="p in product" :key="p.name">
            <td>{{ p.name }}</td>
            <td>{{ p.price *0.01 }}</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
    <div class="text-left">
      <v-chip class="ma-2 px-4" color="primary">
        <v-icon left>mdi-identifier</v-icon>00005f4d9336
      </v-chip>

      <v-chip class="ma-2 px-4" color="secondary">
        <v-icon left>mdi-cellphone-link</v-icon>Android
      </v-chip>

      <v-chip class="ma-2 px-4" color="red" text-color="white">
        <v-icon left>mdi-ip-network</v-icon>99.65.7.60
      </v-chip>

      <v-chip class="ma-2" color="green" text-color="white">
        <v-icon left>mdi-currency-usd</v-icon>248.5
      </v-chip>
    </div>

    <v-simple-table>
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">Name</th>
            <th class="text-left">Price ($ USD)</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="p in product" :key="p.name">
            <td>{{ p.name }}</td>
            <td>{{ p.price *0.01 }}</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
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
          <tr v-for="buyer in relatedIpBuyers" :key="buyer.id">
            <td>{{ buyer.name }}</td>
            <td>{{ buyer.id }}</td>
            <td>{{ buyer.ip }}</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
    <v-divider class="my-4"></v-divider>
    <h2 class="mb-3">Product recomendations</h2>
    <v-simple-table>
      <template v-slot:default>
        <thead>
          <tr>
            <th class="text-left">Name</th>
            <th class="text-left">Price ($ USD)</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="r in recommendations" :key="r.name">
            <td>{{ r.name }}</td>
            <td>{{ r.price * 0.01 }}</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>
  </div>
</template>

<script>
export default {
  data: () => ({
    name: "Sanjiv",
    age: 28,
    product: [
      {
        name: "Chicken base made with roasted chicken",
        price: 9097
      },
      {
        name: "Pepperoni pizza snack rolls",
        price: 5738
      },
      {
        name: "Grilled fajita style carved chicken breast",
        price: 10015
      }
    ],
    relatedIpBuyers: [
      {
        id: "38caa9ba",
        ip: "252.153.172.185",
        name: "Petrine"
      },
      {
        id: "7a48f587",
        ip: "252.153.172.185",
        name: "Oren"
      },
      {
        id: "951348ed",
        ip: "98.94.128.97",
        name: "Koehler"
      },
      {
        id: "79c75bf0",
        ip: "85.237.65.164",
        name: "Allbee"
      },
      {
        id: "da4f6e2a",
        ip: "119.204.182.131",
        name: "Charbonnier"
      }
    ],
    recommendations: [
      {
        name: "Butternut squash & ginger organic soup",
        price: 5336
      },
      {
        name: "Herbalife boisson vanille creme",
        price: 4557
      },
      {
        name: "Campbell's well yes soup vegetable with farro",
        price: 2456
      },
      {
        name: "Progresso Rich & Hearty Chicken & Homestyle Noodles Soup",
        price: 7405
      },
      {
        name: "Combination sausage & pepperoni pizza wrapped in a crispy",
        price: 1698
      },
      {
        name: "Spinach pizza",
        price: 2476
      },
      {
        name: "Organic veggie jalapeno breakfast burrito",
        price: 1605
      },
      {
        name:
          "Twisted hawaiian fresh rising crust pizza topped with rich italian-style sauce",
        price: 2725
      },
      {
        name: "Ranch and bacon pasta salad kit",
        price: 1319
      },
      {
        name: "Long grain & wild rice with herbs & seasonings",
        price: 6405
      }
    ]
  })
};
</script>