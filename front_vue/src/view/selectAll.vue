<template>
    <div>
        <template>
            <el-table
                  :data="tableData"
                  style="width: 60%;margin-top:30px;">
                <el-table-column
                      prop="id"
                      label="ID"
                      width="180">
                </el-table-column>
                <el-table-column
                      prop="name"
                      label="姓名"
                      width="180">
                </el-table-column>
                <el-table-column
                      prop="age"
                      label="年龄">
                </el-table-column>
                <el-table-column
                      label="操作">
                    <template slot-scope="scope">
                        <el-button type="primary" size="small" @click="select(scope.row)">查看</el-button>
                        <el-button type="warning" size="small" @click="update(scope.row)">修改</el-button>
                        <el-button type="danger" size="small" @click="remove(scope.row)">删除</el-button>
                    </template>
                </el-table-column>
            </el-table>
        </template>
    </div>
</template>

<script>
import axios from "axios";

export default {
    name: "selectAll",
    data() {
        return {
            tableData: []
        }
    },
    methods: {
        getData() {
            const _this = this;
            axios.get("http://localhost:8081/student/findAll").then(function (resp) {
                _this.tableData = resp.data;
            })
        },
        remove(stu) {
            const _this = this;
            if (confirm("确定删除吗？")) {
                axios.delete("http://localhost:8081/student/remove/" + stu.id).then(function (resp) {
                    if (resp.data == 1) {
                        _this.getData();
                    }
                })
            }
        },
        select(stu) {
            this.$router.push({
                path: "/selectOne",
                query:{
                    id: stu.id
                }
            })
        },
        update(stu) {
            this.$router.push({
                path: "/update",
                query:{
                    id: stu.id
                }
            })
        }
    },
    created() {
        this.getData();
    }
}
</script>

<style scoped>

</style>
