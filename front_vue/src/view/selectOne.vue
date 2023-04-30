<template>
    <div>
        <el-form :model="ruleForm" status-icon label-width="100px" class="demo-ruleForm"
                 style="margin-top:30px;width: 30%;">
            <el-form-item label="ID" prop="pass">
                <el-input type="text" v-model="ruleForm.id"></el-input>
            </el-form-item>
            <el-form-item label="姓名" prop="pass">
                <el-input type="text" v-model="ruleForm.name"></el-input>
            </el-form-item>
            <el-form-item label="年龄" prop="checkPass">
                <el-input type="text" v-model="ruleForm.age"></el-input>
            </el-form-item>
        </el-form>
    </div>
</template>

<script>
import axios from "axios";

export default {
    name: "selectOne",
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
