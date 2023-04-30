package org.ymx.sp_back.dao.javaMapper;

import org.apache.ibatis.annotations.*;
import org.springframework.stereotype.Repository;
import org.ymx.sp_back.pojo.Student;

import java.util.List;

/**
 * 注解版
 */
@Mapper
@Repository
public interface JavaStudentMapper {

    /**
     * 添加一个学生
     *
     * @param student
     * @return
     */
    @Insert("insert into student(name, age) " +
            "values (#{name}, #{age})")
    public int saveStudent(Student student);

    /**
     * 根据ID查看一名学生
     *
     * @param id
     * @return
     */
    @Select("select *  " +
            "from student " +
            "where id = #{id}")
    public Student findStudentById(Integer id);

    /**
     * 查询全部学生
     *
     * @return
     */
    @Select("select * " +
            "from student")
    @Results({
            @Result(property = "id", column = "id"),
            @Result(property = "name", column = "name"),
            @Result(property = "age", column = "age")
    })
    public List<Student> findAllStudent();

    /**
     * 根据ID删除一个
     *
     * @param id
     * @return
     */
    @Delete("delete   " +
            "from student " +
            "where id = #{id}")
    public int removeStudentById(Integer id);

    /**
     * 根据ID修改
     *
     * @param student
     * @return
     */
    @Update("update set name=#{name},age=#{age}  " +
            "where id=#{id}")
    public int updateStudentById(Student student);

}
