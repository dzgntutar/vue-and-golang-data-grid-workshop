<template>
  <div class="container">
    <product-list-app/>
  </div>
</template>

<script>
import productList from "@/components/product/productList";
export default {
  name: "home",
  components: {
    productListApp:productList
  },created() {
    console.log("Home created")
  }

}
</script>

<style scoped>

</style>