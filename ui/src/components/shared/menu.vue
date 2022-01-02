<template>
  <ul class="nav">

    <li class="nav-item">
      <router-link
          to="/"
          exact
          tag="a"
          class="nav-link"
          active-class="active"
      >Product List
      </router-link
      >

    </li>

    <li class="nav-item">
      <router-link
          to="/addProduct"
          exact
          tag="a"
          class="nav-link"
          active-class="active"
      >Add Product
      </router-link
      >
    </li>

  </ul>
</template>

<script>
export default {
}
</script>

<style scoped>

</style>