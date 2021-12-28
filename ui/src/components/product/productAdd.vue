<template>
  <div class="card myCard">
    <div class="card-header">
      <h4>Product Save</h4>
    </div>
    <div class="card-body">
      <div class="row">
        <div class="col">
          <input type="text" class="form-control" placeholder="Name" v-model="product.name">
        </div>
        <div class="col">
          <input type="number" class="form-control" placeholder="Price" v-model="product.price">
        </div>
      </div>
      <br>
      <div class="row">
        <div class="col">
          <input type="text" class="form-control" placeholder="Count" v-model="product.count">
        </div>
        <div class="col">
          <input type="text" class="form-control" placeholder="Category" v-model="product.category">
        </div>
      </div>
      <br>
      <div class="row">
        <div class="col">
          <button @click="add">Add</button>
        </div>
      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: "productAdd",
  data() {
    return {
      product: {
        name: null,
        price: 0,
        count: 0,
        category: 0
      }
    }
  },
  methods:{
    add(){
      this.$store.dispatch("addProduct", this.product);
    }
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