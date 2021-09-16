<template>
  <el-row>
    <!--顶部-->
    <div class="header">
      <!--顶部左侧-->
      <el-col :span="4" class="logo-container">
        <!--LOGO替代物-->
        <img class="logo" src="../../assets/img/logo.png" alt="logo">
      </el-col>

      <el-col :span="12" class="nav-container">
      </el-col>
      <el-col :span="4" class="user-container" v-if="haveNoToken()">
        <a class="user-button" @click="toLogin">登录</a>
        <a class="user-button" @click="toSignup">注册</a>
      </el-col>
      <el-col :span="4" class="user-container" v-else>
        <div class="user-status">
          <div class="user-avatar" @click="toggleUserPanel">
            <el-collapse-transition>
              <div class="user-panel" v-show="showUserPanel">
                <!--这里可用v-for配合@click使用-->
                <div class="user-function" style="color: #409EFF; font-weight: bold; font-size: 16px">个人信息</div>
                <div class="user-function">用户名：{{ userName }}</div>
                <div class="user-function">用户单位：{{ unit }}</div>
                <div class="user-function" @click="logout"><i class="el-icon-switch-button"></i> 注销</div>
              </div>
            </el-collapse-transition>
          </div>
          <div class="user-info">
            <div class="user-name">{{ userName }}</div>
            <div class="user-id">身份: {{ userRole }}</div>
          </div>
        </div>
      </el-col>
    </div>
  </el-row>
</template>

<script>
export default {
  name: "Header",
  data() {
    return {
      userName: '',
      userId: '',
      unit: '',
      userRole: '',
      showUserPanel: false,
    }
  },
  methods: {
    toggleUserPanel() {
      this.showUserPanel = !this.showUserPanel;
    },
    //检测是否存在token，存在返回false
    haveNoToken() {
      return window.localStorage.getItem('token') === null;
    },
    //注销
    logout() {
      window.localStorage.clear()
      this.$message.success('注销成功!')
      this.$router.replace('/login')
    },
    getUserInfo() {
      this.userName = window.localStorage.getItem('name')
      this.userRole = window.localStorage.getItem('role')
      this.unit = window.localStorage.getItem('unit')
    },
  },
  created() {
    this.getUserInfo()
  }
}
</script>

<style scoped>
* {
  margin: 0;
  padding: 0;
  font-family: "Helvetica Neue", Helvetica, "PingFang SC", "Hiragino Sans GB", "Microsoft YaHei", "微软雅黑", Arial, sans-serif;
  outline: none;
}

.header {
  height: 80px;
  color: white;
  display: flex;
  flex: 1 1 auto;
  justify-content: space-around;
}

.logo-container {
  display: flex;
  justify-content: space-around;
  align-content: center;
}

.logo {
  height: 100%;
  cursor: pointer;
}

.nav-container {
  display: flex;
  /*height: 3rem;*/
}

.user-container {
  display: flex;
  justify-content: center;
  align-items: center;
  /*height: 5rem;*/
}

.user-button {
  color: #409EFF;
  text-decoration-line: none;
  font-size: 0.9rem;
  border: none;
  width: 4rem;
}

.user-status {
  width: 12rem;
  background-color: #FFFFFF11;
  border-radius: 0.5rem;
  padding-top: 0.5rem;
  padding-bottom: 0.5rem;
  display: flex;
  justify-content: space-around;
  align-items: center;
}

.user-avatar {
  position: relative;
  width: 2rem;
  height: 2rem;
  background-color: #2ba250;
  background: url("../../assets/img/logo.png") no-repeat center top;
  background-size: 100%;
  border: 0.05rem solid transparent;
  border-radius: 1rem;
  cursor: pointer;
}

.user-panel {
  color: black;
  text-align: center;
  position: absolute;
  top: 2.5rem;
  left: -0.5rem;
  width: 12rem;
  height: 7.8rem;
  border-radius: 0 0 0.5rem 0.5rem;
  background-color: #ffffff;
  z-index: 888;
}

.user-function {
  line-height: 1.9rem;
  height: 1.9rem;
  border-top: 0.1rem solid #d9d9d9;
}

.user-function:hover {
  background-color: #69c2fc;
}

.user-avatar:hover {
  border: 0.05rem solid white;
}

.user-info {
  cursor: pointer;
  display: flex;
  flex-direction: column;
  justify-content: center;
}

.user-name {
  width: 8rem;
  text-align: center;
  font-weight: bold;
}

.user-name:hover {
  text-decoration-line: underline;
}

.user-id {
  width: 100%;
  text-align: center;
  font-size: 0.8rem;
}

a {
  text-decoration: none;
  color: #ffffff;
  display: inline-block;
  width: 145px;
  height: 80px;
  line-height: 80px;
  text-align: center;
  font-size: 22px;
}

a:hover {
  color: #00fff9;
}
</style>
