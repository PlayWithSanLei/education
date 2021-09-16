<template>
  <div class="container">
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>地图查询</el-breadcrumb-item>
    </el-breadcrumb>
      <iframe src="http://localhost:9427/edu"  class="map"></iframe>
  </div>


</template>

<script>
export default {
  name: "Map"
}
</script>

<style scoped>

.container {
  height: 100%;
}

.map {
  margin: 1em 0 0 0;
  height: 95%;
  width: 100%;
  border: #688383 dashed 1px;
}
</style>
