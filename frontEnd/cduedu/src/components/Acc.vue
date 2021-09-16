<template>

  <div class="container">
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>权限管理</el-breadcrumb-item>
    </el-breadcrumb>

    <div class="left">
      <el-card class="card">
        <el-table></el-table>
      </el-card>
    </div>
    <div class="right">
      <el-card class="card">
        <el-table></el-table>
      </el-card>
    </div>
  </div>
</template>

<script>
export default {
name: "Acc"
}
</script>

<style scoped>

.container {
  height: 100%;
}

.left {
  margin: 1em 3em 0 0;
  height: 98%;
  width: 45%;
  display: inline-block;
}
.right {
  height: 98%;
  width: 45%;
  display: inline-block;
}

.card {
  margin: 1em 0 0 0;
  box-shadow: -1px 1px 20px grey;
  height: 95%;
}
</style>
