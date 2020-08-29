<template>
  <div>
    <div class="search-term">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
        <el-form-item label="类型">
          <el-select clearable placeholder="请选择" v-model="searchInfo.type">
            <el-option
              :key="item.value"
              :label="`${item.label}`"
              :value="item.value"
              v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <el-button @click="onSubmit" type="primary">查询</el-button>
        </el-form-item>
        <el-form-item>
          <el-button @click="openDialog('add')" type="primary">新增收藏</el-button>
        </el-form-item>
      </el-form>
    </div>
    <el-table :data="tableData" @sort-change="sortChange" border stripe>
      <el-table-column label="类型" min-width="50" prop="type" sortable="custom"></el-table-column>
      <el-table-column label="内容" min-width="200" sortable="custom">
        <template slot-scope="scope">
          <a :href=scope.row.url target="_blank" style="color:#606266">
            {{scope.row.title}}
          </a>
        </template>
      </el-table-column>
      <el-table-column label="作者" min-width="80" prop="author" sortable="custom"></el-table-column>
      <el-table-column fixed="right" label="操作" width="200">
        <template slot-scope="scope">
          <el-button @click="editApi(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>
          <el-button @click="deleteApi(scope.row)" size="small" type="danger" icon="el-icon-delete">删除</el-button>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination
      :current-page="page"
      :page-size="pageSize"
      :page-sizes="[20, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="apiForm">
        <el-form-item label="类型" prop="type">
          <el-select placeholder="请选择" v-model="form.type">
            <el-option
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
              v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="内容" prop="type">
          <el-input autocomplete="off" v-model="form.title"></el-input>
        </el-form-item>
        <el-form-item label="链接" prop="apiGroup">
          <el-input autocomplete="off" v-model="form.url"></el-input>
        </el-form-item>
        <el-form-item label="作者" prop="description">
          <el-input autocomplete="off" v-model="form.author"></el-input>
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
// 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

import {
  getCollectionList,
  getCollectionById,
  createCollection,
  updateCollection,
  deleteCollection,
} from "@/api/collection";

import infoList from '@/components/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
const methodOptions = [
  {
    value: 'A用户',
    label: '用户',
    type: ''
  },
  {
    value: 'B专栏',
    label: '专栏',
    type: ''
  },
  {
    value: 'C文章',
    label: '文章',
    type: ''
  },
  {
    value: 'D回答',
    label: '回答',
    type: ''
  },
  {
    value: 'E话术',
    label: '话术',
    type: ''
  }
]

export default {
  name: 'Collection',
  mixins: [infoList],
  data() {
    return {
      listApi: getCollectionList,
      dialogFormVisible: false,
      dialogTitle: '新增收藏',
      form: {
        type: '',
        title: '',
        url: '',
        author: '',
        remark: ''
      },
      methodOptions: methodOptions,
      type: '',
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
  methods: {
    // 排序
    sortChange({ prop, order }) {
      if (prop) {
        this.searchInfo.orderKey = toSQLLine(prop)
        this.searchInfo.desc = order == 'descending'
      }
      this.getTableData()
    },
    //条件搜索前端看此方法
    onSubmit() {
      this.page = 1
      this.pageSize = 20
      this.getTableData()
    },
    initForm() {
      this.$refs.apiForm.resetFields()
      this.form= {
        type: '',
        title: '',
        url: '',
        author: '',
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
      this.type = type
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
            this.getTableData()
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
          switch (this.type) {
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
                this.getTableData()
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
                this.getTableData()
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
  },
  filters: {
    methodFiletr(value) {
      const target = methodOptions.filter(item => item.value === value)[0]
      // return target && `${target.label}(${target.value})`
      return target && `${target.label}`
    },
    tagTypeFiletr(value) {
      const target = methodOptions.filter(item => item.value === value)[0]
      return target && `${target.type}`
    }
  },
  created(){
    this.getTableData()
  }
}
</script>
<style scoped lang="scss">
.button-box {
  padding: 10px 20px;
  .el-button {
    float: right;
  }
}
.el-tag--mini {
  margin-left: 5px;
}
.warning {
  color: #dc143c;
}
</style>