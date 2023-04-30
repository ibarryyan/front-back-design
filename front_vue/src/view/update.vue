<template>
    <div>
        <el-form :model="ruleForm" status-icon label-width="100px" class="demo-ruleForm"
                 style="margin-top:30px;width: 30%;">
            <el-form-item label="ID">
                <el-input type="text" v-model="ruleForm.id" disabled></el-input>
            </el-form-item>
            <el-form-item label="姓名">
                <el-input type="text" v-model="ruleForm.name"></el-input>
            </el-form-item>
            <el-form-item label="年龄">
                <el-input type="text" v-model="ruleForm.age"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="warning" @click="submitForm()">修改</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
import axios from "axios";

export default {
    name: "update",
    data() {
        return {
            ruleForm: {
                id: '',
                name: '',
                age: ''
            }
        };
    },
    methods: {
        submitForm() {
            this.ruleForm.age = this.ruleForm.age - 0
            axios.post("http://localhost:8081/student/update", this.ruleForm).then(function (resp) {
                if (resp.data === 1) {
                    alert("修改成功!")
                } else {
                    alert("修改失败!")
                }
            })
        },
        getStudent() {
            const _this = this;
            axios.get("http://localhost:8081/student/findById/" + this.$route.query.id).then(function (resp) {
                _this.ruleForm = resp.data;
            })
        }
    },
    created() {
        this.getStudent();
    }
}
</script>

<style scoped>

</style>
