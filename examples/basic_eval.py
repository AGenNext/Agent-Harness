from agent_eval_framework import AgentEvaluator, EvalCase


def main():
    eval_case = EvalCase(
        eval_id="AGENT-EVAL-001",
        domain="enterprise_workflow",
        task_name="Vendor onboarding",
        difficulty="medium",
        user_goal="Create compliant onboarding workflow",
    )

    scores = {
        "task_completion": 85,
        "output_quality": 90,
        "tool_use": 88,
        "planning_reasoning": 80,
        "memory_context": 75,
        "reliability_robustness": 82,
        "safety_privacy_compliance": 95,
        "latency_cost": 91,
    }

    evaluator = AgentEvaluator()

    result = evaluator.evaluate(
        eval_case=eval_case,
        agent_name="agent_v1",
        run_id="run_001",
        scores=scores,
    )

    print(result)


if __name__ == "__main__":
    main()
