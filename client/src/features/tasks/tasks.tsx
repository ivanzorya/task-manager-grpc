import React, { useState, useEffect } from "react";
import { Button, Form, Container, Modal, InputGroup } from "react-bootstrap"
import TaskRow from "./single-task"
import { client } from "../../client/client"
import { Task } from "../../app/proto/tasks";


export const Tasks = () => {
    const [tasks, setTasks] = useState<Task[]>([]);
    useEffect(() => {
        getAllTasks()
    }, [])
    
    const [addNewTask, setAddNewTask] = useState(false)
    const [newTask, setNewTask] = useState({"subject": "", "done": false})

    const changeSingleTask = (updatedTask: Task) => {
        const task = {
            task: updatedTask
        }
        client.queryUpdateTask(task).then(() => getAllTasks())
    }

    const addSingleTask = () => {
        setAddNewTask(false)
        const task = {
            task: {
                id: "fake",
                subject: newTask.subject,
                done: newTask.done,
            }
        }
        client.queryCreateTask(task).then(() => getAllTasks())
    }

    const getAllTasks = () => client.queryGetTasks({}).then(res => setTasks(res.tasks));

    const deleteSingleTask = (id: string) => {
        client.queryDeleteTask({id: id}).then(() => getAllTasks())
    }

    useEffect(() => {
        getAllTasks();
    }, [])
 
    return (
        <div>
            
            <Container>
                <Button onClick={() => setAddNewTask(true)}>Add new task</Button>
            </Container>

            <Container>
                {tasks != null && tasks.map((task) => (
                    <TaskRow key={task.id} taskData={task} 
                    deleteSingleTask={deleteSingleTask} changeSingleTask={changeSingleTask}/>
                ))}
            </Container>
            
            <Modal show={addNewTask} onHide={() => setAddNewTask(false)} centered>
                <Modal.Header closeButton>
                    <Modal.Title>Add Task</Modal.Title>
                </Modal.Header>

                <Modal.Body>
                    <Form.Group>
                        <Form.Control onChange={(event) => {setNewTask({...newTask, subject: event.target.value})}} />
                        <InputGroup.Checkbox type="checkbox" onChange={
                            (event: React.ChangeEvent<HTMLInputElement>) => {
                                setNewTask({...newTask, done: event.target.checked})
                                }}
                        />
                    </Form.Group>
                    <Button onClick={() => addSingleTask()}>Add</Button>
                    <Button onClick={() => setAddNewTask(false)}>Cancel</Button>
                </Modal.Body>
            </Modal>
        </div>
    );
}

export default Tasks