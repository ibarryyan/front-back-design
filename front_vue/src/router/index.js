//注册路由
import Vue from 'vue';
import VueRouter from 'vue-router';
//引入路由
import index from '../view/index'
import update from "../view/update";
import selectAll from "../view/selectAll";
import selectOne from "../view/selectOne";
import insert from "../view/insert";

Vue.use(VueRouter);

const router = new VueRouter({
    routes: [
        {
            name: "主页重定向",
            path: "/",
            redirect: "/index"
        }, {
            name: "主页",
            path: "/index",
            component: index,
            children: [
                {
                    name: "修改操作",
                    path: "/update",
                    component: update,
                }, {
                    name: "查看全部",
                    path: "/selectAll",
                    component: selectAll,
                }, {
                    name: "查看一个",
                    path: "/selectOne",
                    component: selectOne,
                }, {
                    name: "添加一个",
                    path: "/insert",
                    component: insert,
                }
            ]
        }
    ]
})

export default router


