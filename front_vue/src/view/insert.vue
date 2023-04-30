<template>
    <div>
        <el-form :model="ruleForm" status-icon label-width="100px" class="demo-ruleForm"
                 style="margin-top:30px;width: 30%;">
            <el-form-item label="姓名" prop="pass">
                <el-input type="text" v-model="ruleForm.name"></el-input>
            </el-form-item>
            <el-form-item label="年龄" prop="checkPass">
                <el-input type="text" v-model="ruleForm.age"></el-input>
            </el-form-item>
            <el-form-item>
                <el-button type="primary" @click="submitForm()">提交</el-button>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
import axios from 'axios'

export default {
    name: "insert",
    data() {
        return {
            ruleForm: {
                name: '',
                age: ''
            }
        };
    },
    methods: {
        submitForm() {
            const _this = this
            this.ruleForm.age = this.ruleForm.age - 0 //转为int
            axios.post("http://localhost:8081/student/save", this.ruleForm).then(function (resp) {
                if (resp.data === 1) {
                    alert("添加成功!")
                    _this.ruleForm.name = ''
                    _this.ruleForm.age = ''
                } else {
                    alert("添加失败!")
                }
            })
        },
    }
}
</script>

<style scoped>

</style>
