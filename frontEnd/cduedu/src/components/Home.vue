<template>
  <el-container class="container">
    <el-header class="header" height="80px">
      <Header></Header>
    </el-header>
    <el-container>
      <el-aside width="200px" class="aside">
        <el-menu class="menu"
                 :default-active="activePath"
                 background-color="#3a3d49"
                 text-color="white"
                 active-text-color="#419DFA"
                 unique-opened
                 router
        >
          <el-sub-menu index="/user" class="el-sub-menu">
            <template #title>
              <i class="el-icon-user-solid"></i>
              <span>组织机构</span>
            </template>
            <el-menu-item index="/user" @click="saveNavState('/user')">
              <template #title>
                <i class="el-icon-s-custom"></i>
                <span>用户列表</span>
              </template>
            </el-menu-item>
            <el-menu-item index="/org" @click="saveNavState('/org')">
              <template #title>
                <i class="el-icon-menu"></i>
                <span>组织列表</span>
              </template>
            </el-menu-item>
          </el-sub-menu>
          <el-menu-item index="/acc" @click="saveNavState('/Acc')" class="el-sub-menu">
            <template #title>
              <i class="el-icon-key"></i>
              <span>权限管理</span>
            </template>
          </el-menu-item>
          <el-menu-item index="/map" class="el-sub-menu" @click="saveNavState('/map')">
            <i class="el-icon-search"></i>
            <template #title>地图查询</template>
          </el-menu-item>
        </el-menu>
      </el-aside>
      <el-main class="main">
        <router-view></router-view>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
import Header from "@/components/public/Header";

export default {
  name: "Home",
  components: {
    Header
  },
  data() {
    return {
      activePath: ''
    }
  },
  methods: {
    saveNavState(path) {
      window.localStorage.setItem('activePath', path)
    },
    getNavState() {
      this.activePath = window.localStorage.getItem('activePath')
    }
  },
  created() {
    this.getNavState()
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
  overflow: auto;
}

.container {
  height: 100%;
}


.menu {
  font-weight: bold;
}

.el-sub-menu {
  width: 200px;
}
</style>
