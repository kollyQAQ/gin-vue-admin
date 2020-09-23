<template>
    <div>
        <div class="search-term">
            <el-form :inline="true" :model="searchInfo" class="demo-form-inline">
                <el-form-item label="类型">
                    <el-select clearable placeholder="请选择" v-model="searchInfo.type" type="number">
                        <el-option value=1 label="正常日志"></el-option>
                        <el-option value=2 label="错误日志"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item label="事件">
                    <el-select clearable placeholder="请选择" v-model="searchInfo.event">
                        <el-option value="sync_goods" label="sync_goods"></el-option>
                        <el-option value="sync_question" label="sync_question"></el-option>
                        <el-option value="sync_answer" label="sync_answer"></el-option>
                        <el-option value="sync_article_rank" label="sync_article_rank"></el-option>
                    </el-select>
                </el-form-item>
                <el-form-item>
                    <el-button @click="onSubmit" type="primary">查询</el-button>
                </el-form-item>
            </el-form>
        </div>
        <el-table :data="tableData" @sort-change="sortChange" border stripe>
            <el-table-column label="id" min-width="80" prop="ID"></el-table-column>
            <el-table-column label="类型" min-width="100" prop="type">
                <template slot-scope="scope">
                    <p v-if="scope.row.type == 1">
                        <el-tag type="success">正常日志</el-tag>
                    </p>
                    <p v-if="scope.row.type == 2">
                        <el-tag type="danger">错误日志</el-tag>
                    </p>
                </template>
            </el-table-column>
            <el-table-column label="事件" min-width="150" prop="event"></el-table-column>
            <el-table-column label="描述" min-width="500" prop="msg"></el-table-column>
            <el-table-column label="日志时间" min-width="200" prop="update_time"></el-table-column>
<!--            <el-table-column fixed="right" label="操作" width="200">-->
<!--                <template slot-scope="scope">-->
<!--                    <el-button @click="editGoods(scope.row)" size="small" type="primary" icon="el-icon-edit">编辑</el-button>-->
<!--                    <el-button @click="deleteGoods(scope.row)" size="small" type="danger" icon="el-icon-delete">删除</el-button>-->
<!--                </template>-->
<!--            </el-table-column>-->
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
    </div>
</template>


<script>
    // 获取列表内容封装在mixins内部  getTableData方法 初始化已封装完成 条件搜索时候 请把条件安好后台定制的结构体字段 放到 this.searchInfo 中即可实现条件搜索

    import {
        queryLog,
    } from '@/api/stat'

    import infoList from '@/components/mixins/infoList'

    export default {
        name: 'Log',
        mixins: [infoList],
        data() {
            return {
                listApi: queryLog,
            }
        },
        methods: {
            onSubmit() {
                this.page = 1
                this.pageSize = 20
                this.getTableData()
            },
        },

        async created(){
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
