<template>
  <div class="card myCard">
    <div class="card-header">
      <h4>Product List</h4>
    </div>
    <div class="card-body">
      <table class="table">
        <thead>
        <tr>
          <th scope="col">#</th>
          <th scope="col">Name</th>
          <th scope="col">Price</th>
          <th scope="col">Count</th>
          <th scope="col">Category</th>
        </tr>
        </thead>
        <tbody>
        <tr
            v-for="product in getAllProducts"
            :key="product.name"
          >

          <th scope="row">{{ product.id }}</th>
          <td>{{ product.name }}</td>
          <td>{{ product.price }}</td>
          <td>{{product.count}}</td>
          <td>{{product.category}}</td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import {mapGetters} from "vuex"
export default {

  name: "productList",
  computed:{
    ...mapGetters(["getAllProducts"])
  }

}
</script>

<style scoped>
.myCard {
  width: 50% !important;
  display: grid;
  /*place-content: center;*/
  margin-left: auto;
  margin-right: auto;
}
</style>