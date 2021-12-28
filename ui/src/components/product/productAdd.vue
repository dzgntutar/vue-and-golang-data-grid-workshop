<template>
  <div class="card myCard">
    <div class="card-header">
      <h4>Product Save</h4>
    </div>
    <div class="card-body">
      <div class="row">
        <div class="col">
          <input type="text" class="form-control" placeholder="Name">
        </div>
        <div class="col">
          <input type="number" class="form-control" placeholder="Price">
        </div>
      </div>
      <br>
      <div class="row">
        <div class="col">
          <input type="text" class="form-control" placeholder="Count">
        </div>
        <div class="col">
          <input type="text" class="form-control" placeholder="Category">
        </div>
      </div>
      <br>
      <div class="row">
        <div class="col">
          <button>Add</button>
        </div>

      </div>
    </div>

  </div>
</template>

<script>
export default {
  name: "productAdd"
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