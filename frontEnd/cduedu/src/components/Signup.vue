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
        <el-button class="button" type="primary" @click="submitForm()">注册</el-button>
        <el-button class="button" @click="resetForm('ruleForm')">重置</el-button>
      </el-form>
    </el-card>
  </div>
</template>

<script>
import {isEmail} from "@/assets/JS/validateEmail";
// import axios from "axios";
export default {
  name: "Signup",
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
    let validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'))
      } else if (value !== this.ruleForm.password) {
        callback(new Error('两次输入密码不一致!'))
      } else {
        callback()
      }
    };
    let validateEmail = (rule, value, callback) => {
      if(!isEmail(value)) {
        callback(new Error('邮箱格式错误！'))
      } else {
        callback()
      }
    }

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
        repassword: [{required: true, validator: validatePass2, trigger: 'blur'}],
        username: [{required: true, message: '请输入姓名', trigger: 'blur'}],
        question: [{required: true, message: '请选择密保问题', trigger: 'blur'}],
        answer: [{required: true, message: '请输入密保答案', trigger: 'blur'}],
        email: [{required: true, message: '请输入邮箱', trigger: 'blur'},{validator: validateEmail, trigger: 'blur'}],
        mobile: [{required: true, message: '请输入手机号', trigger: 'blur'}, {validator: validateMobile, trigger: 'blur'}]
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
      options: ['你的父亲名字', '你的母亲名字', '你的宠物名字', '你的大学名字', '你的名字']
    }
  },
  methods: {
    resetForm(formName) {
      this.$refs[formName].resetFields()
    },
    async submitForm() {
      const {data:res} = await this.$axios.post('/signup', this.ruleForm)
      // console.log(res)
      if (res.code === 1000) {
        this.$router.replace('/login')
        this.$message.success('注册成功')
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
  height: 40em;
  box-shadow: 1px -1px 2px #666666;
  border: 30px solid rgba(255, 255, 255, 0.4);
  background-clip: padding-box;
}

.question {
  width: 100%;
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
</style>
