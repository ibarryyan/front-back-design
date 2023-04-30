package org.ymx.sp_back.service.impl;

import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;
import org.ymx.sp_back.dao.xmlMapper.XmlStudentMapper;
import org.ymx.sp_back.pojo.Student;
import org.ymx.sp_back.service.StudentService;

import java.util.List;

/**
 * @author 17122
 */
@Service
public class StudentServiceImpl implements StudentService {

    @Autowired
    private XmlStudentMapper xmlStudentMapper;

    @Override
    public int saveStudent(Student student) {
        return xmlStudentMapper.saveStudent(student);
    }

    @Override
    public Student findStudentById(Integer id) {
        return xmlStudentMapper.findStudentById(id);
    }

    @Override
    public List<Student> findAllStudent() {
        return xmlStudentMapper.findAllStudent();
    }

    @Override
    public int removeStudentById(Integer id) {
        return xmlStudentMapper.removeStudentById(id);
    }

    @Override
    public int updateStudentById(Student student) {
        return xmlStudentMapper.updateStudentById(student);
    }
}
