package org.ymx.sp_back.controller;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;
import org.ymx.sp_back.pojo.Student;
import org.ymx.sp_back.service.StudentService;

import java.util.List;

/**
 * @author 17122
 */
@RestController
@RequestMapping("/student")
public class StudentController {

    @Autowired
    private StudentService studentService;

    /**
     * 添加学生
     *
     * @param student
     * @return
     */
    @PostMapping("/save")
    public int saveStudent(@RequestBody Student student) {
        int result;
        try {
            result = studentService.saveStudent(student);
        } catch (Exception exception) {
            return -1;
        }
        return result;
    }

    /**
     * 查看全部
     *
     * @return
     */
    @GetMapping("/findAll")
    public List<Student> findAll() {
        return studentService.findAllStudent();
    }

    /**
     * 根据ID查看
     *
     * @param id
     * @return
     */
    @GetMapping("/findById/{id}")
    public Student findById(@PathVariable("id") Integer id) {
        return studentService.findStudentById(id);
    }

    /**
     * 删除一个
     *
     * @param id
     * @return
     */
    @DeleteMapping("/remove/{id}")
    public int remove(@PathVariable("id") Integer id) {
        return studentService.removeStudentById(id);
    }

    /**
     * 修改学生信息
     *
     * @param student
     * @return
     */
    @PostMapping("/update")
    public int update(@RequestBody Student student) {
        return studentService.updateStudentById(student);
    }

}
