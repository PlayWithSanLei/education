<template>
  <div class="container">
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>组织机构</el-breadcrumb-item>
      <el-breadcrumb-item>组织列表</el-breadcrumb-item>
    </el-breadcrumb>
    <el-card class="card">
      <el-table></el-table>
    </el-card>
  </div>

</template>

<script>
export default {
  name: "Org"
}
</script>

<style scoped>

.container {
  height: 100%;
}

.card {
  margin: 1em 0 0 0;
  box-shadow: -1px 1px 20px grey;
  height: 95%;
}
</style>
