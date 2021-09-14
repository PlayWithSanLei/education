<template>
  <div class="container">
    <el-card class="login-panel">
      <h2 class="title">注册您的账户</h2>
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
        <el-form-item label="姓名" prop="username">
          <el-input
              type="text"
              v-model="ruleForm.username"
              autocomplete="off"></el-input>
        </el-form-item>
        <el-form-item label="邮箱" prop="email">
          <el-input type="text" v-model="ruleForm.email" autocomplete="off"></el-input>
        </el-form-item>

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
        <el-form-item label="确认密码" prop="repassword">
          <el-input
              type="password"
              v-model="ruleForm.repassword"
              autocomplete="off"
          ></el-input>
        </el-form-item>
        <el-form-item label="密保问题" prop="question">
          <el-select class="question" v-model="ruleForm.question" clearable placeholder="请选择密保问题">
            <el-option v-for="item in options" :key="item" :label="item" :value="item">
            </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="密保答案" prop="answer">
          <el-input type="text" v-model="ruleForm.answer" autocomplete="off"></el-input>
        </el-form-item>
        <br/>
        <el-button class="button" type="primary" @click="submitForm()">提交</el-button>
        <el-button class="button" @click="resetForm('ruleForm')">重置</el-button>
      </el-form>
    </el-card>
  </div>
</template>

<script>
export default {
  name: "Signup",
  data() {
    var validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'))
      } else {
        if (this.ruleForm.repassword !== '') {
          this.$refs.ruleForm.validateField('repassword')
        }
        callback()
      }
    }
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.ruleForm.password) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    }
    return {
      rules: {
        password: [{validator: validatePass, trigger: 'blur'}],
        repassword: [{validator: validatePass2, trigger: 'blur'}],
      },
      ruleForm: {
        username: '',
        email: '',
        mobile: '',
        password: '',
        repassword: '',
        question: '',
        answer: ''
      },
      options: ['你的父亲名字','你的母亲名字','你的宠物名字','你的大学名字','你的名字']
    }
  },
  methods: {
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    submitForm() {
      console.log(JSON.parse(JSON.stringify(this.ruleForm)))
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
  height: 40em;
  box-shadow: 1px -1px 2px #666666;
  border: 30px solid rgba(255, 255, 255, 0.4);
  background-clip: padding-box;
}
.question {
  width: 29.8em;
}
.linkTop {
  width: 100%;
  height: 1px;
  border-top:solid #bbbbbb 1px;
}
.demo-ruleForm {
  padding-right: 2.7em;
}
.button {
  float: right;
  margin-left: 1em;
}
</style>
