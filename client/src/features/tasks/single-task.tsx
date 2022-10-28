import React, { useState } from "react";
import { Button, Card, Row, Col, InputGroup } from "react-bootstrap";
import { Task } from "../../app/proto/tasks";


const TaskRow: React.FC<{
    taskData: Task, 
    deleteSingleTask: (id: string) => void,
    changeSingleTask: (task: Task) => void
  }> = ({taskData, deleteSingleTask, changeSingleTask}) => {
    const [done, setDone] = useState(taskData.done)

    return (
        <Card>
            <Row>
                <Col>Subject:{ taskData !== undefined && taskData.subject}</Col>
                <Col><InputGroup.Checkbox checked={done} onChange={(event: React.ChangeEvent<HTMLInputElement>) => {
                    setDone(event.target.checked);
                    changeSingleTask({id: taskData.id, subject:taskData.subject, done: event.target.checked});
                    }}
                /></Col>
                <Col><Button onClick={() => deleteSingleTask(taskData.id)}>delete task</Button></Col>
            </Row>
        </Card>
    )
}

export default TaskRow