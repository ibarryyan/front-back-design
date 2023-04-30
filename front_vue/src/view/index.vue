<template>
    <div>
        <el-menu class="el-menu-demo" mode="horizontal" :router="true">
            <el-menu-item index="/selectAll">全部学生</el-menu-item>
            <el-menu-item index="/insert">添加</el-menu-item>
            <el-menu-item index="/selectOne" disabled>查看</el-menu-item>
            <el-menu-item index="/update" disabled>修改</el-menu-item>
        </el-menu>
        <router-view></router-view>
    </div>
</template>

<script>
export default {
    name: "index"
}
</script>

<style scoped>

</style>
