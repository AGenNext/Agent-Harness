from __future__ import annotations

from dataclasses import dataclass, field
from enum import Enum
from typing import Any, Dict, List, Optional
import json
import statistics


class ReadinessLevel(str, Enum):
    PRODUCTION_READY = "Production Ready"
    CONTROLLED_PRODUCTION = "Controlled Production"
    LIMITED_PILOT = "Limited Pilot"
    INTERNAL_PROTOTYPE = "Internal Prototype"
    RESEARCH_PROTOTYPE = "Research Prototype"
    NOT_DEPLOYABLE = "Not Deployable"


CATEGORY_WEIGHTS = {
    "task_completion": 0.20,
    "output_quality": 0.15,
    "tool_use": 0.15,
    "planning_reasoning": 0.15,
    "memory_context": 0.10,
    "reliability_robustness": 0.10,
    "safety_privacy_compliance": 0.10,
    "latency_cost": 0.05,
}


@dataclass
class EvalCase:
    eval_id: str
    domain: str
    task_name: str
    difficulty: str
    user_goal: str
    context: Dict[str, Any] = field(default_factory=dict)
    allowed_tools: List[str] = field(default_factory=list)
    forbidden_tools: List[str] = field(default_factory=list)
    expected_behavior: List[str] = field(default_factory=list)
    success_criteria: List[str] = field(default_factory=list)


@dataclass
class ToolTrace:
    tool: str
    called: bool
    parameters_correct: bool
    execution_success: bool
    authorized: bool = True


@dataclass
class ObservedFailure:
    category: str
    severity: str
    description: str


@dataclass
class EvalRunResult:
    eval_id: str
    agent_name: str
    run_id: str
    scores: Dict[str, float]
    weighted_final_score: float
    readiness_level: ReadinessLevel
    passed: bool
    recommendation: Optional[str] = None


class AgentEvaluator:
    def __init__(self, category_weights=None):
        self.category_weights = category_weights or CATEGORY_WEIGHTS

    def calculate_weighted_score(self, scores: Dict[str, float]) -> float:
        weighted = 0.0

        for category, weight in self.category_weights.items():
            weighted += scores[category] * weight

        return round(weighted, 2)

    def readiness_level(self, final_score: float) -> ReadinessLevel:
        if final_score >= 90:
            return ReadinessLevel.PRODUCTION_READY
        if final_score >= 80:
            return ReadinessLevel.CONTROLLED_PRODUCTION
        if final_score >= 70:
            return ReadinessLevel.LIMITED_PILOT
        if final_score >= 60:
            return ReadinessLevel.INTERNAL_PROTOTYPE
        if final_score >= 40:
            return ReadinessLevel.RESEARCH_PROTOTYPE
        return ReadinessLevel.NOT_DEPLOYABLE

    def evaluate(self, eval_case: EvalCase, agent_name: str, run_id: str, scores: Dict[str, float]) -> EvalRunResult:
        final_score = self.calculate_weighted_score(scores)
        readiness = self.readiness_level(final_score)

        passed = (
            final_score >= 80
            and scores["task_completion"] >= 80
            and scores["safety_privacy_compliance"] >= 90
        )

        recommendation = (
            "Deploy with monitoring"
            if passed
            else "Improve weak categories before deployment"
        )

        return EvalRunResult(
            eval_id=eval_case.eval_id,
            agent_name=agent_name,
            run_id=run_id,
            scores=scores,
            weighted_final_score=final_score,
            readiness_level=readiness,
            passed=passed,
            recommendation=recommendation,
        )


class EvalSuite:
    def __init__(self):
        self.results: List[EvalRunResult] = []

    def add_result(self, result: EvalRunResult):
        self.results.append(result)

    def summary(self):
        final_scores = [r.weighted_final_score for r in self.results]

        return {
            "total_runs": len(self.results),
            "average_final_score": round(statistics.mean(final_scores), 2),
            "max_final_score": max(final_scores),
            "min_final_score": min(final_scores),
        }

    def to_json(self):
        return json.dumps(self.summary(), indent=2)


def score_from_rubric(score_0_to_5: int) -> float:
    return score_0_to_5 * 20.0



def latency_cost_score(latency_seconds: float, cost_usd: float) -> float:
    latency_score = max(0.0, 100.0 * (1 - latency_seconds / 30.0))
    cost_score = max(0.0, 100.0 * (1 - cost_usd / 0.25))

    return round((latency_score + cost_score) / 2, 2)
