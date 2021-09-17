<template>
  <div class="container">
    <el-breadcrumb separator-class="el-icon-arrow-right">
      <el-breadcrumb-item :to="{ path: '/home' }">首页</el-breadcrumb-item>
      <el-breadcrumb-item>组织机构</el-breadcrumb-item>
      <el-breadcrumb-item>用户列表</el-breadcrumb-item>
    </el-breadcrumb>
    <el-card class="card">
      <el-row :gutter="20">
        <el-col :span="7">
          <el-input placeholder="请输入内容" v-model="queryPara.value" class="input-with-select" clearable
                    @clear="getUserInfo">
            <template #prepend>
              <el-select v-model="queryPara.column" placeholder="请选择查询字段" class="el-select">
                <el-option key="用户id" value="userid" label="用户id"></el-option>
                <el-option key="用户名" value="username" label="用户名"></el-option>
                <el-option key="手机号" value="mobile" label="手机号"></el-option>
                <el-option key="邮箱" value="email" label="邮箱"></el-option>
                <el-option key="单位" value="unit" label="单位"></el-option>
                <el-option key="角色" value="role" label="角色"></el-option>
              </el-select>
            </template>
            <template #append>
              <el-button icon="el-icon-search"></el-button>
            </template>
          </el-input>
        </el-col>
        <el-col :span="4">
        </el-col>
        <el-col :span="9">
        </el-col>
      </el-row>
      <!--        列表区域-->
      <el-table :data="userlist.slice((currPage-1)*PageSize,currPage*PageSize)" border stripe class="table">
        <el-table-column type="index" label="#"></el-table-column>
        <el-table-column label="用户id" prop="userid"></el-table-column>
        <el-table-column label="用户名" prop="username"></el-table-column>
        <el-table-column label="手机号" prop="mobile"></el-table-column>
        <el-table-column label="邮箱" prop="email"></el-table-column>
        <el-table-column label="单位" prop="unit"></el-table-column>
        <el-table-column label="角色" prop="role"></el-table-column>
        <el-table-column label="用户状态" prop="status" width="125px">
          <template #default="scope">
            <el-switch v-model="scope.row.status" @change="editUserStatus(scope.row.userid)" active-color="#13ce66"
                       inactive-color="#ff4949">
            </el-switch>
          </template>
        </el-table-column>
        <el-table-column label="操作" width="125px">
          <template #default="scope">
            <!--              修改-->
            <el-tooltip effect="dark" content="修改用户" placement="top" :enterable="false">
              <el-button type="primary" icon="el-icon-edit" size="mini"
                         @click="showEditDialog(scope.row)"></el-button>
            </el-tooltip>
            <!--              删除-->
            <el-tooltip effect="dark" content="删除用户" placement="top" :enterable="false">
              <el-button type="danger" icon="el-icon-delete" size="mini"
                         @click="remove(scope.row.userid)"></el-button>
            </el-tooltip>
          </template>
        </el-table-column>
      </el-table>

      <!--              分页区域-->
      <el-pagination
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="currPage"
          :page-sizes="[10, 30]"
          :page-size="PageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
      >
      </el-pagination>

      <!--      dialog-->
      <el-dialog title="修改用户" v-model="editDialogVisible" class="edit-dialog">
        <el-form :model="editForm">
          <el-form-item label="用户id" :label-width="100" class="edit-input">
            <el-input v-model="editForm.userid" autocomplete="off" disabled></el-input>
          </el-form-item>
          <el-form-item label="用户名称" :label-width="100" class="edit-input">
            <el-input v-model="editForm.username" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="手机号" :label-width="100" class="edit-input">
            <el-input v-model="editForm.mobile" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="邮箱" :label-width="100" class="edit-input">
            <el-input v-model="editForm.email" autocomplete="off"></el-input>
          </el-form-item>
          <el-form-item label="单位" :label-width="100">
            <el-select v-model="editForm.unit" placeholder="请选择单位" class="edit-select">
              <el-option label="区域一" value="shanghai"></el-option>
              <el-option label="区域二" value="beijing"></el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="角色" :label-width="100">
            <el-select v-model="editForm.role" placeholder="请选择角色" class="edit-select">
              <el-option label="区域一" value="shanghai"></el-option>
              <el-option label="区域二" value="beijing"></el-option>
            </el-select>
          </el-form-item>
        </el-form>
        <template #footer>
    <span class="dialog-footer">
      <el-button @click="editDialogVisible = false">取 消</el-button>
      <el-button type="primary" @click="editDialogVisible = false"
      >确 定</el-button
      >
    </span>
        </template>
      </el-dialog>
      <el-dialog title="删除用户" v-model="removeDialogVisible" width="30%" @close="getUserInfo">
        <span style="font-size: 18px">确定删除这个用户吗?</span>
        <template #footer>
        <span class="dialog-footer">
        <el-button @click="cancelEdit()">取 消</el-button>
        <el-button type="primary" @click="removeUserById()">确 定</el-button>
      </span>
        </template>
      </el-dialog>
      <el-table></el-table>
    </el-card>
  </div>

</template>

<script>
export default {
  name: "User",
  data() {
    return {
      currPage: 1,
      PageSize: 10,
      queryPara: {
        value: '',
        column: '',
      },
      userlist: [],
      rolesList: [],
      units: [],
      total: 0,
      removeDialogVisible: false,
      editDialogVisible: false,
      editForm: {
        userid: '',
        username: '',
        mobile: '',
        email: '',
        unit: '',
        role: '',
        status: '',
      },
      selectedRoleId: '',
      removeID: ''
    }
  },
  methods: {
    handleSizeChange(newSize) {
      this.currPage = 1
      this.PageSize = newSize
    },
    // 监听页码改变
    handleCurrentChange(newPage) {
      this.currPage = newPage
    },
    async getUserInfo() {
      const {data: res} = await this.$axios.get('/home/users')
      console.log(res)
      console.log(res.msg)
      this.userlist = res.msg
      this.total = res.msg.length
    },
    showEditDialog(row) {
      this.editDialogVisible = true
      this.editForm = row
    },
    editUserStatus(id) {
      this.$axios.head('/home/users/' + id)
    },
    remove(id) {
      this.removeDialogVisible = true
      this.removeID = id
    },
    removeUserById() {
      this.$axios.delete('/home/users/' + this.removeID)
      this.removeDialogVisible = false
      this.getUserInfo()
    },
    cancelEdit() {
      this.removeDialogVisible = false
      this.getUserInfo()
    }
  },
  created() {
    this.getUserInfo()
  }
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

.table {
  margin-top: 2em;
}

.el-select {
  width: 11em;
}

.input-with-select {
  width: 50em;
}

.edit-dialog {
  display: flex;
  justify-content: center;
  align-items: center;
}

.edit-input {
  width: 90%;
}

.edit-select {
  width: 88%;
  position: relative;
  left: -6%;
}
</style>
