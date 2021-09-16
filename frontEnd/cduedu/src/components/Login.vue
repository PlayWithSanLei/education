<template>
  <div class="container">
    <el-card class="login-panel">
      <h2 class="title">民办教育信息服务与监管平台 <span class="version">v1.0</span> </h2>
      <span class="login">登录</span>
      <div class="linkTop"></div>
      <br>
      <el-form
          :model="ruleForm"
          status-icon
          :rules="rules"
          ref="ruleForm"
          label-width="100px"
          class="demo-ruleForm"
      >
        <el-form-item label="手机号" prop="mobile">
          <el-input type="text" v-model="ruleForm.mobile" autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="密码" prop="password">
          <el-input
              type="password"
              v-model="ruleForm.password"
              autocomplete="off"
          ></el-input>
        </el-form-item>
        <br/>
        <el-button class="button" type="primary" @click="submitForm()">登录</el-button>
        <el-button class="button" @click="goToSignup()">注册</el-button>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "Login",
  data() {
    let validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'))
      } else {
        if (this.ruleForm.repassword !== '') {
          this.$refs.ruleForm.validateField('repassword')
        }
        callback()
      }
    };

    let validateMobile = (rule, value, callback) => {
      if(!(/^[1]{1}[0-9]{10}$/.test(value))) {
        callback(new Error('请输入正确的11位手机号码！'))
      }
      else {
        callback()
      }
    }
    return {
      rules: {
        password: [{required: true,  validator: validatePass, trigger: 'blur'}],
        mobile: [{required: true, message: '请输入手机号', trigger: 'blur'}, {validator: validateMobile, trigger: 'blur'}]
      },
      ruleForm: {
        mobile: '',
        password: '',
      },
    }
  },
  methods: {
    goToSignup() {
      this.$router.replace('/signup')
    },
    async submitForm() {
      const {data:res} = await this.$axios.post('/login', this.ruleForm)
      console.log(res)
      if (res.code === 1000) {
        let token = 'Bearer '  + res.msg.token
        window.localStorage.setItem('token', token)
        window.localStorage.setItem('name', res.msg.name)
        window.localStorage.setItem('role', res.msg.role)
        window.localStorage.setItem('unit', res.msg.unit)
        this.$message.success('登录成功')
        this.$router.replace('/home')
      } else {
        this.$message.error('用户名或密码错误')
      }
    }
  }
}
</script>

<style scoped>
.container {
  background: url("../assets/img/bg.jpeg");
  height: 100%;
  position: relative;
  display: flex;
  justify-content: center;
  align-items: center;
}

.login-panel {
  width: 600px;
  height: 22em;
  box-shadow: 1px -1px 2px #666666;
  border: 30px solid rgba(255, 255, 255, 0.4);
  background-clip: padding-box;
}

.title {
  color: #2881dc;
}

.linkTop {
  width: 100%;
  height: 1px;
  border-top: solid #bbbbbb 1px;
}

.demo-ruleForm {
  padding-right: 2.7em;
}

.button {
  float: right;
  margin-left: 1em;
}

.login {
  font-weight: bold;
  font-size: 1.5em;
  position: relative;
  left: -9em
}

.version {
  float: right;
  font-weight: lighter;
  font-size: 0.8em;
  margin: 0.5em 0 0 0;
}
</style>
