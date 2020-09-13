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
          <tr v-for="item in desserts" :key="item.name">
            <td>{{ item.name }}</td>
            <td>{{ item.age }}</td>
            <td>{{ item.id }}</td>
          </tr>
        </tbody>
      </template>
    </v-simple-table>

    <div class="text-center">
      <v-btn class="ma-5" dark>Load more</v-btn>
    </div>
  </div>
</template>

<script>
export default {
  data() {
    return {
      desserts: [
        {
          id: "aaf0f8ab",
          name: "Rella",
          age: 65
        },
        {
          id: "e612df58",
          name: "Ermentrude",
          age: 30
        },
        {
          id: "1be7fe04",
          name: "McDermott",
          age: 62
        },
        {
          id: "1c8ecf7c",
          name: "Dolora",
          age: 62
        },
        {
          id: "83e45e62",
          name: "Mayfield",
          age: 62
        },
        {
          id: "64534d5e",
          name: "Hardie",
          age: 40
        },
        {
          id: "cd6a86b3",
          name: "Paul",
          age: 35
        },
        {
          id: "591408e5",
          name: "Elurd",
          age: 77
        },
        {
          id: "72978269",
          name: "Idel",
          age: 21
        },
        {
          id: "b119e66e",
          name: "Kissel",
          age: 48
        }
      ]
    };
  }
};
</script>
