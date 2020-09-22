<template>
  <div>
    <el-form :inline="true" class="demo-form-inline">
      <el-form-item>
        <el-button @click="openDialog('add')" type="primary">新增收藏</el-button>
      </el-form-item>
    </el-form>

    <el-table :data="list" border fit highlight-current-row style="width: 100%">
      <el-table-column
              v-loading="loading"
              align="center"
              label="ID"
              width="65"
              element-loading-text="请给我点时间！"
      >
        <template slot-scope="scope">
          <span>{{ scope.row.ID }}</span>
        </template>
      </el-table-column>

      <el-table-column min-width="250px" label="Title">
        <template slot-scope="{row}">
          <a :href=row.url target="_blank" style="color:#606266">
            {{row.title}}
          </a>
        </template>
      </el-table-column>

      <el-table-column width="150px" align="center" label="Author">
        <template slot-scope="scope">
          <span>{{ scope.row.author }}</span>
        </template>
      </el-table-column>

      <el-table-column label="优先级" min-width="80" prop="priority" ></el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editApi(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-button @click="deleteApi(scope.row)" size="small" type="danger" icon="el-icon-delete">删除</el-button>
        </template>
      </el-table-column>
    </el-table>

    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="apiForm">
        <el-form-item label="类型" prop="type">
          <el-select placeholder="请选择" v-model="form.type">
            <el-option
                    :key="item.value"
                    :label="`${item.label}`"
                    :value="item.value"
                    v-for="item in typeOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="内容">
          <el-input autocomplete="off" v-model="form.title"></el-input>
        </el-form-item>
        <el-form-item label="链接">
          <el-input autocomplete="off" v-model="form.url"></el-input>
        </el-form-item>
        <el-form-item label="作者">
          <el-input autocomplete="off" v-model="form.author"></el-input>
        </el-form-item>
        <el-form-item label="优先级">
          <el-input autocomplete="off" v-model="form.priority"></el-input>
        </el-form-item>
      </el-form>
      <div class="edit_container">
        <quill-editor
                :options="{}"
                ref="myQuillEditor"
                v-model="form.remark"
        ></quill-editor>
      </div>
      <div class="dialog-footer" slot="footer">
        <el-button @click="closeDialog">取 消</el-button>
        <el-button @click="enterDialog" type="primary">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
import {
  getCollectionListV2,
  getCollectionById,
  createCollection,
  updateCollection,
  deleteCollection,
} from "@/api/collection";

export default {
  props: {
    type: {
      type: String,
      default: '文章'
    }
  },
  data() {
    return {
      list: null,
      listQuery: {
        type: this.type,
        sort: '+id'
      },
      loading: false,

      dialogFormVisible: false,
      dialogTitle: '新增收藏',
      form: {
        type: this.type,
        title: '',
        url: '',
        author: '',
        priority: '0',
        remark: ''
      },
      typeOptions: [
        {
          value: '用户',
          label: '用户',
        },
        {
          value: '专栏',
          label: '专栏',
        },
        {
          value: '文章',
          label: '文章',
        },
        {
          value: '回答',
          label: '回答',
        },
        {
          value: '话术',
          label: '话术',
        }
      ],
      rules: {
        type: [{ required: true, message: '请输入type', trigger: 'blur' }],
        title: [
          { required: true, message: '请输入title', trigger: 'blur' }
        ],
        url: [
          { required: false, message: '请选择url', trigger: 'blur' }
        ],
        author: [
          { required: false, message: '请输入author', trigger: 'blur' }
        ]
      }
    }
  },
  created() {
    this.getList()
  },
  methods: {
    getList() {
      this.loading = true
      this.$emit('create') // for test
      getCollectionListV2(this.listQuery).then(response => {
        this.list = response.data.list
        this.loading = false
      })
    },
    initForm() {
      this.$refs.apiForm.resetFields()
      this.form= {
        type: this.type,
        title: '',
        url: '',
        author: '',
        priority: '0',
        remark: ''
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    openDialog(type) {
      switch (type) {
        case 'add':
          this.dialogTitle = '新增收藏'
          break
        case 'edit':
          this.dialogTitle = '编辑收藏'
          break
        default:
          break
      }
      this.dialogType = type
      this.dialogFormVisible = true
    },
    async editApi(row) {
      const res = await getCollectionById({ id: row.ID })
      this.form = res.data.collection
      this.openDialog('edit')
    },
    async deleteApi(row) {
      this.$confirm('此操作将删除此收藏, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
              .then(async () => {
                const res = await deleteCollection(row)
                if (res.code == 0) {
                  this.$message({
                    type: 'success',
                    message: '删除成功!'
                  })
                  this.getList()
                }
              })
              .catch(() => {
                this.$message({
                  type: 'info',
                  message: '已取消删除'
                })
              })
    },
    async enterDialog() {
      this.$refs.apiForm.validate(async valid => {
        if (valid) {
          switch (this.dialogType) {
            case 'add':
            {
              const res = await createCollection(this.form)
              if (res.code == 0) {
                this.$message({
                  type: 'success',
                  message: '添加成功',
                  showClose: true
                })
              }
              this.getList()
              this.closeDialog()
            }

              break
            case 'edit':
            {
              const res = await updateCollection(this.form)
              if (res.code == 0) {
                this.$message({
                  type: 'success',
                  message: '编辑成功',
                  showClose: true
                })
              }
              this.getList()
              this.closeDialog()
            }
              break
            default:
            {
              this.$message({
                type: 'error',
                message: '未知操作',
                showClose: true
              })
            }
              break
          }
        }
      })
    }
  }
}
</script>

