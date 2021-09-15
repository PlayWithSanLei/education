<template>
  <el-container class="container">
    <el-header class="header" height="80px"><Header></Header></el-header>
    <el-container>
      <el-aside width="200px" class="aside"></el-aside>
      <el-main class="main"></el-main>
    </el-container>
  </el-container>
</template>

<script>
import Header from "@/components/public/Header";
export default {
  name: "Home",
  components: {
    Header
  }
}
</script>

<style scoped>
.header {
  background-color: #222222;
}

.aside {
  background-color: #3A3D49;

}

.main {
  background-color: #EBEDF1;
}

.container {
  height: 100%;
}
</style>
