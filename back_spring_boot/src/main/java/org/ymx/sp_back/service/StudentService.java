package org.ymx.sp_back.service;

import org.ymx.sp_back.pojo.Student;

import java.util.List;

/**
 * @author 17122
 */
public interface StudentService {
    /**
     * 添加一个学生
     *
     * @param student
     * @return
     */
    public int saveStudent(Student student);

    /**
     * 根据ID查看一名学生
     *
     * @param id
     * @return
     */
    public Student findStudentById(Integer id);

    /**
     * 查询全部学生
     *
     * @return
     */
    public List<Student> findAllStudent();

    /**
     * 根据ID删除一个
     *
     * @param id
     * @return
     */
    public int removeStudentById(Integer id);

    /**
     * 根据ID修改
     *
     * @param student
     * @return
     */
    public int updateStudentById(Student student);


}
