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
          <el-button @click="openDialog('addApi')" type="primary">新增商品</el-button>
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
      :page-sizes="[30, 50, 100]"
      :style="{float:'right',padding:'20px'}"
      :total="total"
      @current-change="handleCurrentChange"
      @size-change="handleSizeChange"
      layout="total, sizes, prev, pager, next, jumper"
    ></el-pagination>

    <el-dialog :before-close="closeDialog" :title="dialogTitle" :visible.sync="dialogFormVisible">
      <el-form :inline="true" :model="form" :rules="rules" label-width="80px" ref="apiForm">
        <el-form-item label="路径" prop="path">
          <el-input autocomplete="off" v-model="form.path"></el-input>
        </el-form-item>
        <el-form-item label="请求" prop="method">
          <el-select placeholder="请选择" v-model="form.method">
            <el-option
              :key="item.value"
              :label="`${item.label}(${item.value})`"
              :value="item.value"
              v-for="item in methodOptions"
            ></el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="api分组" prop="apiGroup">
          <el-input autocomplete="off" v-model="form.apiGroup"></el-input>
        </el-form-item>
        <el-form-item label="api简介" prop="description">
          <el-input autocomplete="off" v-model="form.description"></el-input>
        </el-form-item>
      </el-form>
      <div class="warning">新增Api需要在角色管理内配置权限才可使用</div>
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
  getApiById,
  createApi,
  updateApi,
  deleteApi
} from '@/api/api'

import {
  getCollectionList
} from "../../../api/collection";

import infoList from '@/components/mixins/infoList'
import { toSQLLine } from '@/utils/stringFun'
const methodOptions = [
  {
    value: '用户',
    label: '用户',
    type: ''
  },
  {
    value: '专栏',
    label: '专栏',
    type: ''
  },
  {
    value: '文章',
    label: '文章',
    type: ''
  },
  {
    value: '回答',
    label: '回答',
    type: ''
  },
  {
    value: '话术',
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
        path: '',
        apiGroup: '',
        method: '',
        description: ''
      },
      methodOptions: methodOptions,
      type: '',
      rules: {
        path: [{ required: true, message: '请输入api路径', trigger: 'blur' }],
        apiGroup: [
          { required: true, message: '请输入组名称', trigger: 'blur' }
        ],
        method: [
          { required: true, message: '请选择请求方式', trigger: 'blur' }
        ],
        description: [
          { required: true, message: '请输入api介绍', trigger: 'blur' }
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
        sku_id: '',
        sku_name: '',
        third_category: '',
        fee_rate: -1
      }
    },
    closeDialog() {
      this.initForm()
      this.dialogFormVisible = false
    },
    openDialog(type) {
      switch (type) {
        case 'addApi':
          this.dialogTitlethis = '新增Api'
          break
        case 'edit':
          this.dialogTitlethis = '编辑Api'
          break
        default:
          break
      }
      this.type = type
      this.dialogFormVisible = true
    },
    async editApi(row) {
      const res = await getApiById({ id: row.ID })
      this.form = res.data.api
      this.openDialog('edit')
    },
    async deleteApi(row) {
      this.$confirm('此操作将永久删除所有角色下该菜单, 是否继续?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
      })
        .then(async () => {
          const res = await deleteApi(row)
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
            case 'addApi':
              {
                const res = await createApi(this.form)
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
                const res = await updateApi(this.form)
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