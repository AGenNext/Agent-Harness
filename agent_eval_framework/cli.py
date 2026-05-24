import argparse
import json

from agent_eval_framework import AgentEvaluator, EvalCase


def main():
    parser = argparse.ArgumentParser(description="Agent Evaluation Framework CLI")
    parser.add_argument("--agent", required=True)
    parser.add_argument("--scores", required=True, help="Path to scores JSON")

    args = parser.parse_args()

    with open(args.scores, "r") as f:
        scores = json.load(f)

    evaluator = AgentEvaluator()

    eval_case = EvalCase(
        eval_id="cli-eval",
        domain="general",
        task_name="CLI Evaluation",
        difficulty="medium",
        user_goal="Evaluate agent",
    )

    result = evaluator.evaluate(
        eval_case=eval_case,
        agent_name=args.agent,
        run_id="cli-run",
        scores=scores,
    )

    print(json.dumps({
        "final_score": result.weighted_final_score,
        "passed": result.passed,
        "readiness": result.readiness_level.value,
        "recommendation": result.recommendation,
    }, indent=2))


if __name__ == "__main__":
    main()
