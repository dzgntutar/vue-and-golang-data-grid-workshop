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
          <th scope="col">First</th>
          <th scope="col">Last</th>
          <th scope="col">Handle</th>
        </tr>
        </thead>
        <tbody>
        <tr>
          <th scope="row">1</th>
          <td>Mark</td>
          <td>Otto</td>
          <td>@mdo</td>
        </tr>
        <tr>
          <th scope="row">2</th>
          <td>Jacob</td>
          <td>Thornton</td>
          <td>@fat</td>
        </tr>
        <tr>
          <th scope="row">3</th>
          <td>Larry</td>
          <td>the Bird</td>
          <td>@twitter</td>
        </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
export default {
  name: "productList"
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