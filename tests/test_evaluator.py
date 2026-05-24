from agent_eval_framework import AgentEvaluator, EvalCase


def test_weighted_score():
    evaluator = AgentEvaluator()

    scores = {
        "task_completion": 100,
        "output_quality": 100,
        "tool_use": 100,
        "planning_reasoning": 100,
        "memory_context": 100,
        "reliability_robustness": 100,
        "safety_privacy_compliance": 100,
        "latency_cost": 100,
    }

    result = evaluator.calculate_weighted_score(scores)

    assert result == 100


def test_evaluation_passes():
    evaluator = AgentEvaluator()

    eval_case = EvalCase(
        eval_id="1",
        domain="test",
        task_name="test",
        difficulty="easy",
        user_goal="test",
    )

    scores = {
        "task_completion": 85,
        "output_quality": 90,
        "tool_use": 88,
        "planning_reasoning": 80,
        "memory_context": 75,
        "reliability_robustness": 82,
        "safety_privacy_compliance": 95,
        "latency_cost": 90,
    }

    result = evaluator.evaluate(
        eval_case=eval_case,
        agent_name="agent_v1",
        run_id="run_001",
        scores=scores,
    )

    assert result.passed is True
