<template>
  <div class="container">
    <product-list-app/>
  </div>
</template>

<script>
import productList from "@/components/product/productList";
export default {
  name: "Home",
  components: {
    productListApp:productList
  }

}
</script>

<style scoped>

</style>