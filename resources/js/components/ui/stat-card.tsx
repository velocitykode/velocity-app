import { type LucideIcon } from 'lucide-react';

type ColorVariant = 'blue' | 'violet' | 'emerald' | 'amber' | 'rose';

interface StatCardProps {
  /** Stat label */
  label: string;
  /** Stat value */
  value: string | number;
  /** Icon component */
  icon: LucideIcon;
  /** Color theme */
  color?: ColorVariant;
  /** Optional change indicator (e.g., "+2") */
  change?: string;
  /** Optional status text */
  status?: string;
  /** Additional classes */
  className?: string;
}

const colorClasses: Record<ColorVariant, string> = {
  blue: 'bg-blue-100 text-[#1e3a8a] dark:bg-blue-900/50 dark:text-blue-400',
  violet: 'bg-violet-100 text-violet-600 dark:bg-violet-900/50 dark:text-violet-400',
  emerald: 'bg-emerald-100 text-emerald-600 dark:bg-emerald-900/50 dark:text-emerald-400',
  amber: 'bg-amber-100 text-amber-600 dark:bg-amber-900/50 dark:text-amber-400',
  rose: 'bg-rose-100 text-rose-600 dark:bg-rose-900/50 dark:text-rose-400',
};

export function StatCard({
  label,
  value,
  icon: Icon,
  color = 'blue',
  change,
  className = '',
}: StatCardProps) {
  return (
    <div
      className={`rounded-2xl border border-slate-200 bg-white p-4 lg:p-6 transition-all duration-200 hover:shadow-lg hover:-translate-y-0.5 dark:border-zinc-800 dark:bg-zinc-900 ${className}`}
    >
      <div className={`inline-flex rounded-xl p-2 lg:p-2.5 ${colorClasses[color]}`}>
        <Icon className="h-4 w-4 lg:h-5 lg:w-5" />
      </div>
      <div className="mt-3 lg:mt-4 flex items-baseline gap-1 lg:gap-2">
        <span className="text-2xl lg:text-3xl font-bold text-slate-900 dark:text-white">
          {value}
        </span>
        {change && (
          <span className="text-xs lg:text-sm font-medium text-emerald-600 dark:text-emerald-400">
            {change}
          </span>
        )}
      </div>
      <p className="mt-1 text-xs lg:text-sm text-slate-500 dark:text-zinc-400">{label}</p>
    </div>
  );
}
